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
	g.GET("", h.UserRentals)
	g.GET("/:id", h.RentalDetail)
	g.POST("/create", h.CreateRental)
	g.POST("/cancel", h.CancelRental)
}

func (h *RentalHandler) UserRentals(c echo.Context) error {
	userID := c.Get("user_id").(int)

	rentals, err := h.rentalUsecase.UserRentals(c.Request().Context(), userID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get rentals success", rentals)
}

func (h *RentalHandler) RentalDetail(c echo.Context) error {
	userID := c.Get("user_id").(int)
	rentalID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid rental id")
	}

	rental, err := h.rentalUsecase.RentalDetail(c.Request().Context(), userID, rentalID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get rental success", rental)
}

func (h *RentalHandler) CreateRental(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateRentalRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	rental, err := h.rentalUsecase.CreateRental(c.Request().Context(), userID, &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, "create rental success", rental)
}

func (h *RentalHandler) CancelRental(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CancelRentalRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	if err := h.rentalUsecase.CancelRental(c.Request().Context(), userID, req.RentalID); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "cancel rental success", nil)
}
