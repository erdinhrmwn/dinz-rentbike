package contract

import (
	"context"

	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type ReviewRepository interface {
	FindByUserID(ctx context.Context, userID int) ([]entity.Review, error)
	FindByRentalID(ctx context.Context, rentalID int) (*entity.Review, error)
	FindByVehicleID(ctx context.Context, vehicleID int) ([]entity.Review, error)
	Create(ctx context.Context, review *entity.Review) error
	Update(ctx context.Context, review *entity.Review) error
	Delete(ctx context.Context, id int) error
}

type ReviewUsecase interface {
	GetByUserID(ctx context.Context, userID int) ([]dto.ReviewResponse, error)
	Create(ctx context.Context, userID int, req *dto.CreateReviewRequest) (*dto.ReviewResponse, error)
	Update(ctx context.Context, userID int, req *dto.CreateReviewRequest) (*dto.ReviewResponse, error)
	Delete(ctx context.Context, userID int, rentalID int) error
}
