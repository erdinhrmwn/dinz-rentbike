package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/pkg/response"
)

type AdminPaymentHandler struct {
	paymentUsecase contract.PaymentUsecase
}

func NewAdminPaymentHandler(paymentUsecase contract.PaymentUsecase) *AdminPaymentHandler {
	return &AdminPaymentHandler{paymentUsecase: paymentUsecase}
}

func (h *AdminPaymentHandler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.List)
	g.GET("/:id", h.Detail)
}

func (h *AdminPaymentHandler) List(c echo.Context) error {
	payments, err := h.paymentUsecase.ListAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get payments success", payments)
}

func (h *AdminPaymentHandler) Detail(c echo.Context) error {
	paymentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid payment id")
	}

	payment, err := h.paymentUsecase.AdminDetail(c.Request().Context(), paymentID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get payment success", payment)
}
