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
	g.GET("", h.UserReviews)
	g.POST("/create", h.CreateReview)
	g.POST("/edit", h.UpdateReview)
	g.POST("/delete", h.DeleteReview)
}

func (h *ReviewHandler) UserReviews(c echo.Context) error {
	userID := c.Get("user_id").(int)

	reviews, err := h.reviewUsecase.UserReviews(c.Request().Context(), userID)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "get reviews success", reviews)
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateReviewRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	review, err := h.reviewUsecase.CreateReview(c.Request().Context(), userID, &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, "create review success", review)
}

func (h *ReviewHandler) UpdateReview(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateReviewRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	review, err := h.reviewUsecase.UpdateReview(c.Request().Context(), userID, &req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "update review success", review)
}

func (h *ReviewHandler) DeleteReview(c echo.Context) error {
	userID := c.Get("user_id").(int)

	var req dto.CreateReviewRequest
	if err := c.Bind(&req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	if err := h.reviewUsecase.DeleteReview(c.Request().Context(), userID, req.RentalID); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "delete review success", nil)
}
