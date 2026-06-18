package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/pkg/response"
)

type AdminVehicleHandler struct {
	vehicleUsecase contract.VehicleUsecase
}

func NewAdminVehicleHandler(vehicleUsecase contract.VehicleUsecase) *AdminVehicleHandler {
	return &AdminVehicleHandler{vehicleUsecase: vehicleUsecase}
}

func (h *AdminVehicleHandler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.Create)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *AdminVehicleHandler) Create(c echo.Context) error {
	var req dto.CreateVehicleRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	vehicle, err := h.vehicleUsecase.Create(c.Request().Context(), &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, "create vehicle success", vehicle)
}

func (h *AdminVehicleHandler) Update(c echo.Context) error {
	vehicleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid vehicle id")
	}

	var req dto.UpdateVehicleRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	vehicle, err := h.vehicleUsecase.Update(c.Request().Context(), vehicleID, &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "update vehicle success", vehicle)
}

func (h *AdminVehicleHandler) Delete(c echo.Context) error {
	vehicleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid vehicle id")
	}

	if err := h.vehicleUsecase.Delete(c.Request().Context(), vehicleID); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "delete vehicle success", nil)
}
