package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/pkg/response"
)

type AdminRentalHandler struct {
	rentalUsecase contract.RentalUsecase
}

func NewAdminRentalHandler(rentalUsecase contract.RentalUsecase) *AdminRentalHandler {
	return &AdminRentalHandler{rentalUsecase: rentalUsecase}
}

func (h *AdminRentalHandler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.List)
	g.GET("/:id", h.Detail)
	g.PATCH("/:id/status", h.UpdateStatus)
}

func (h *AdminRentalHandler) List(c echo.Context) error {
	rentals, err := h.rentalUsecase.ListAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get rentals success", rentals)
}

func (h *AdminRentalHandler) Detail(c echo.Context) error {
	rentalID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid rental id")
	}

	rental, err := h.rentalUsecase.AdminDetail(c.Request().Context(), rentalID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get rental success", rental)
}

func (h *AdminRentalHandler) UpdateStatus(c echo.Context) error {
	rentalID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid rental id")
	}

	var req dto.UpdateRentalStatusRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	rental, err := h.rentalUsecase.UpdateStatus(c.Request().Context(), rentalID, req.Status)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "update rental status success", rental)
}
