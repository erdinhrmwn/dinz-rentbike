package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/pkg/response"
)

type RentalHandler struct {
	rentalUsecase contract.RentalUsecase
}

func NewRentalHandler(rentalUsecase contract.RentalUsecase) *RentalHandler {
	return &RentalHandler{rentalUsecase: rentalUsecase}
}

func (h *RentalHandler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.GetByUserID)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.PATCH("/:id/cancel", h.Cancel)
}

func (h *RentalHandler) GetByUserID(c echo.Context) error {
	userID := c.Get("user_id").(int)

	rentals, err := h.rentalUsecase.GetByUserID(c.Request().Context(), userID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get rentals success", rentals)
}

func (h *RentalHandler) GetByID(c echo.Context) error {
	userID := c.Get("user_id").(int)
	rentalID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid rental id")
	}

	rental, err := h.rentalUsecase.GetByID(c.Request().Context(), userID, rentalID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get rental success", rental)
}

func (h *RentalHandler) Create(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateRentalRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	rental, err := h.rentalUsecase.Create(c.Request().Context(), userID, &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, "create rental success", rental)
}

func (h *RentalHandler) Cancel(c echo.Context) error {
	userID := c.Get("user_id").(int)
	rentalID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid rental id")
	}

	if err := h.rentalUsecase.Cancel(c.Request().Context(), userID, rentalID); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "cancel rental success", nil)
}
