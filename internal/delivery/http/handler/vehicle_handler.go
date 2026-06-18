package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/pkg/response"
)

type VehicleHandler struct {
	vehicleUsecase contract.VehicleUsecase
}

func NewVehicleHandler(vehicleUsecase contract.VehicleUsecase) *VehicleHandler {
	return &VehicleHandler{vehicleUsecase: vehicleUsecase}
}

func (h *VehicleHandler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
}

func (h *VehicleHandler) GetAll(c echo.Context) error {
	vehicles, err := h.vehicleUsecase.GetAll(c.Request().Context())
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get vehicles success", vehicles)
}

func (h *VehicleHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid vehicle id")
	}

	vehicle, err := h.vehicleUsecase.GetByID(c.Request().Context(), id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get vehicle success", vehicle)
}
