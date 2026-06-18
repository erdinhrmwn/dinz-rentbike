package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/pkg/response"
)

type UserHandler struct {
	userUsecase contract.UserUsecase
}

func NewUserHandler(userUsecase contract.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/me", h.GetProfile)
	g.PUT("/update", h.UpdateProfile)
	g.PATCH("/change-password", h.ChangePassword)
}

func (h *UserHandler) GetProfile(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid user id")
	}

	profile, err := h.userUsecase.GetProfile(c.Request().Context(), userID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusNotFound, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get profile success", profile)
}

func (h *UserHandler) UpdateProfile(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid user id")
	}

	var req dto.UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	if err := h.userUsecase.UpdateProfile(c.Request().Context(), userID, &req); err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "update profile success", nil)
}

func (h *UserHandler) ChangePassword(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid user id")
	}

	var req dto.ChangePasswordRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	if err := h.userUsecase.ChangePassword(c.Request().Context(), userID, &req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "change password success", nil)
}
