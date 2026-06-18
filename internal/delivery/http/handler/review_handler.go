package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/pkg/response"
)

type ReviewHandler struct {
	reviewUsecase contract.ReviewUsecase
}

func NewReviewHandler(reviewUsecase contract.ReviewUsecase) *ReviewHandler {
	return &ReviewHandler{reviewUsecase: reviewUsecase}
}

func (h *ReviewHandler) RegisterRoutes(g *echo.Group) {
	g.GET("", h.GetByUserID)
	g.POST("/create", h.Create)
	g.POST("/edit", h.Update)
	g.POST("/delete", h.Delete)
}

func (h *ReviewHandler) GetByUserID(c echo.Context) error {
	userID := c.Get("user_id").(int)

	reviews, err := h.reviewUsecase.GetByUserID(c.Request().Context(), userID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get reviews success", reviews)
}

func (h *ReviewHandler) Create(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateReviewRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	review, err := h.reviewUsecase.Create(c.Request().Context(), userID, &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, "create review success", review)
}

func (h *ReviewHandler) Update(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateReviewRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	review, err := h.reviewUsecase.Update(c.Request().Context(), userID, &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "update review success", review)
}

func (h *ReviewHandler) Delete(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateReviewRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	if err := h.reviewUsecase.Delete(c.Request().Context(), userID, req.RentalID); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "delete review success", nil)
}
