package handler

import (
	"encoding/json"
	"io"
	"net/http"

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
}

func NewWebhookHandler(webhookToken string, paymentUsecase contract.PaymentUsecase, rentalUsecase contract.RentalUsecase) *WebhookHandler {
	return &WebhookHandler{webhookToken: webhookToken, paymentUsecase: paymentUsecase, rentalUsecase: rentalUsecase}
}

func (h *WebhookHandler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.XenditWebhook)
}

type xenditPaymentSessionWebhookPayload struct {
	Event      string `json:"event"`
	BusinessID string `json:"business_id"`
	Data       struct {
		PaymentSessionID string  `json:"payment_session_id"`
		Status           string  `json:"status"`
		ReferenceID      string  `json:"reference_id"`
		PaymentID        string  `json:"payment_id"`
		Amount           float64 `json:"amount"`
		Currency         string  `json:"currency"`
		Country          string  `json:"country"`
	} `json:"data"`
}

func (h *WebhookHandler) XenditWebhook(c echo.Context) error {
	// Verify callback token
	callbackToken := c.Request().Header.Get("x-callback-token")
	if callbackToken != h.webhookToken {
		logger.Log.Warn().Str("ip", c.RealIP()).Str("callback_token", callbackToken).Msg("xendit webhook: invalid callback token")
		return response.ErrorResponse(c, http.StatusUnauthorized, "invalid callback token")
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "failed to read body")
	}

	var payload xenditPaymentSessionWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		logger.Log.Error().Err(err).Str("body", string(body)).Msg("xendit webhook: failed to parse")
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid payload")
	}

	logger.Log.Info().
		Str("event", payload.Event).
		Str("session_id", payload.Data.PaymentSessionID).
		Str("status", payload.Data.Status).
		Msg("xendit webhook received")

	// Find payment by xendit session ID
	payment, err := h.paymentUsecase.FindByInvoiceID(c.Request().Context(), payload.Data.PaymentSessionID)
	if err != nil {
		logger.Log.Warn().Str("session_id", payload.Data.PaymentSessionID).Msg("payment not found for session")
		return response.ErrorResponse(c, http.StatusNotFound, "payment not found")
	}

	// Update payment status based on webhook
	switch payload.Data.Status {
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

	return response.SuccessResponse(c, http.StatusOK, "webhook processed", updatedPayment)
}
