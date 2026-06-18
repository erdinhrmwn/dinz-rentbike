package usecase

import (
	"context"
	"errors"
	"time"

	"dinz-rentbike/internal/domain/constants"
	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type reviewUsecase struct {
	reviewRepo contract.ReviewRepository
	rentalRepo contract.RentalRepository
}

func NewReviewUsecase(reviewRepo contract.ReviewRepository, rentalRepo contract.RentalRepository) contract.ReviewUsecase {
	return &reviewUsecase{reviewRepo: reviewRepo, rentalRepo: rentalRepo}
}

func (u *reviewUsecase) UserReviews(ctx context.Context, userID int) ([]dto.ReviewResponse, error) {
	reviews, err := u.reviewRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var res []dto.ReviewResponse
	for _, r := range reviews {
		res = append(res, toReviewResponse(&r))
	}
	return res, nil
}

func (u *reviewUsecase) CreateReview(ctx context.Context, userID int, req *dto.CreateReviewRequest) (*dto.ReviewResponse, error) {
	rental, err := u.rentalRepo.FindByID(ctx, req.RentalID)
	if err != nil {
		return nil, err
	}

	if rental.UserID != userID {
		return nil, errors.New("rental not found")
	}

	if rental.Status != constants.RentalStatusCompleted {
		return nil, errors.New("only completed rental can be reviewed")
	}

	review := &entity.Review{
		UserID:    userID,
		VehicleID: rental.VehicleID,
		RentalID:  rental.ID,
		Rating:    req.Rating,
	}

	if req.Comment != "" {
		review.Comment = &req.Comment
	}

	if err := u.reviewRepo.Create(ctx, review); err != nil {
		return nil, err
	}

	res := toReviewResponse(review)
	return &res, nil
}

func (u *reviewUsecase) UpdateReview(ctx context.Context, userID int, req *dto.CreateReviewRequest) (*dto.ReviewResponse, error) {
	review, err := u.reviewRepo.FindByRentalID(ctx, req.RentalID)
	if err != nil {
		return nil, err
	}

	if review.UserID != userID {
		return nil, errors.New("review not found")
	}

	review.Rating = req.Rating
	if req.Comment != "" {
		review.Comment = &req.Comment
	}

	if err := u.reviewRepo.Update(ctx, review); err != nil {
		return nil, err
	}

	res := toReviewResponse(review)
	return &res, nil
}

func (u *reviewUsecase) DeleteReview(ctx context.Context, userID int, rentalID int) error {
	review, err := u.reviewRepo.FindByRentalID(ctx, rentalID)
	if err != nil {
		return err
	}

	if review.UserID != userID {
		return errors.New("review not found")
	}

	return u.reviewRepo.Delete(ctx, review.ID)
}

func toReviewResponse(r *entity.Review) dto.ReviewResponse {
	res := dto.ReviewResponse{
		ID:        r.ID,
		UserID:    r.UserID,
		VehicleID: r.VehicleID,
		RentalID:  r.RentalID,
		Rating:    r.Rating,
		Comment:   r.Comment,
		CreatedAt: r.CreatedAt.Format(time.RFC3339),
		UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
	}

	if r.Vehicle != nil {
		res.Vehicle = &dto.VehicleResponse{
			ID:           r.Vehicle.ID,
			Type:         r.Vehicle.Type,
			Brand:        r.Vehicle.Brand,
			Name:         r.Vehicle.Name,
			Category:     r.Vehicle.Category,
			Description:  r.Vehicle.Description,
			ImageURL:     r.Vehicle.ImageURL,
			PricePerHour: r.Vehicle.PricePerHour,
			Status:       r.Vehicle.Status,
			CreatedAt:    r.Vehicle.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    r.Vehicle.UpdatedAt.Format(time.RFC3339),
		}
	}

	return res
}
