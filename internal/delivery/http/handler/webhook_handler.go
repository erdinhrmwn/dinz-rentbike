package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/constants"
	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/pkg/logger"
	"dinz-rentbike/pkg/response"
)

type WebhookHandler struct {
	webhookToken   string
	paymentUsecase contract.PaymentUsecase
	rentalUsecase  contract.RentalUsecase
	mailjetService contract.MailjetService
}

func NewWebhookHandler(webhookToken string, paymentUsecase contract.PaymentUsecase, rentalUsecase contract.RentalUsecase, mailjetService contract.MailjetService) *WebhookHandler {
	return &WebhookHandler{webhookToken: webhookToken, paymentUsecase: paymentUsecase, rentalUsecase: rentalUsecase, mailjetService: mailjetService}
}

func (h *WebhookHandler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.XenditWebhook)
}

type xenditWebhookPayload struct {
	Event      string `json:"event"`
	BusinessID string `json:"business_id"`
	Data       any      `json:"data"`
}

type paymentSessionPayload struct {
	PaymentSessionID string  `json:"payment_session_id"`
	Status           string  `json:"status"`
	ReferenceID      string  `json:"reference_id"`
	PaymentID        string  `json:"payment_id"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	Country          string  `json:"country"`
	Metadata         struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"metadata"`
}

func (h *WebhookHandler) XenditWebhook(c echo.Context) error {
	callbackToken := c.Request().Header.Get("x-callback-token")
	if callbackToken != h.webhookToken {
		logger.Log.Warn().Str("ip", c.RealIP()).Str("callback_token", callbackToken).Msg("xendit webhook: invalid callback token")
		return response.ErrorResponse(c, http.StatusUnauthorized, "invalid callback token")
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "failed to read body")
	}

	var payload xenditWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		logger.Log.Error().Err(err).Str("body", string(body)).Msg("xendit webhook: failed to parse")
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid payload")
	}

	if !strings.HasPrefix("payment_session.", payload.Event) {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid event")
	}

	data, ok := payload.Data.(paymentSessionPayload)
	if !ok {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid data")
	}

	logger.Log.Info().
		Str("event", payload.Event).
		Str("session_id", data.PaymentSessionID).
		Str("status", data.Status).
		Msg("xendit webhook received")

	payment, err := h.paymentUsecase.FindByInvoiceID(c.Request().Context(), data.PaymentSessionID)
	if err != nil {
		logger.Log.Warn().Str("session_id", data.PaymentSessionID).Msg("payment not found for session")
		return response.ErrorResponse(c, http.StatusNotFound, "payment not found")
	}

	switch data.Status {
	case "COMPLETED":
		payment.Status = constants.PaymentStatusPaid
	case "EXPIRED":
		payment.Status = constants.PaymentStatusExpired
	default:
		return response.SuccessResponse(c, http.StatusOK, "webhook received", nil)
	}

	updatedPayment, err := h.paymentUsecase.UpdatePayment(c.Request().Context(), payment.ID, payment.Status)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, "failed to update payment")
	}

	// Send email notification if payment completed
	if updatedPayment.Status == constants.PaymentStatusPaid {
		name := data.Metadata.Name
		go h.mailjetService.Send(&contract.EmailRequest{
			ToEmail: data.Metadata.Email,
			ToName:  name,
			Subject: "Pembayaran Berhasil - Dinz RentBike",
			HTMLBody: fmt.Sprintf(`
				<h2>Pembayaran Berhasil! ✅</h2>
				<p>Halo %s, terima kasih! Pembayaran Anda sebesar <strong>Rp %.0f</strong> telah kami terima.</p>
				<p>Selamat menikmati perjalanan Anda! 🏍️</p>
			`, name, data.Amount),
		})
	}

	return response.SuccessResponse(c, http.StatusOK, "webhook processed", updatedPayment)
}
