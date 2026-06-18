package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/pkg/response"
)

type PaymentHandler struct {
	paymentUsecase contract.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase contract.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUsecase: paymentUsecase}
}

func (h *PaymentHandler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.UserPayments)
	g.GET("/:id", h.PaymentDetail)
	g.POST("", h.CreatePayment)
	g.POST("/cancel", h.CancelPayment)
}

func (h *PaymentHandler) UserPayments(c echo.Context) error {
	userID := c.Get("user_id").(int)

	payments, err := h.paymentUsecase.UserPayments(c.Request().Context(), userID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get payments success", payments)
}

func (h *PaymentHandler) CreatePayment(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreatePaymentRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	payment, err := h.paymentUsecase.CreatePayment(c.Request().Context(), userID, req.RentalID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, "create payment success", payment)
}

func (h *PaymentHandler) CancelPayment(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CancelPaymentRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	payment, err := h.paymentUsecase.CancelPayment(c.Request().Context(), userID, req.RentalID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "cancel payment success", payment)
}

func (h *PaymentHandler) PaymentDetail(c echo.Context) error {
	userID := c.Get("user_id").(int)
	paymentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid payment id")
	}

	payment, err := h.paymentUsecase.PaymentDetail(c.Request().Context(), userID, paymentID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get payment success", payment)
}
