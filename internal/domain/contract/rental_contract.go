package contract

import (
	"context"

	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type RentalRepository interface {
	FindByID(ctx context.Context, id int) (*entity.Rental, error)
	FindByUserID(ctx context.Context, userID int) ([]entity.Rental, error)
	Create(ctx context.Context, rental *entity.Rental) error
	Update(ctx context.Context, rental *entity.Rental) error
}

type RentalUsecase interface {
	GetByID(ctx context.Context, userID int, rentalID int) (*dto.RentalResponse, error)
	GetByUserID(ctx context.Context, userID int) ([]dto.RentalResponse, error)
	Create(ctx context.Context, userID int, req *dto.CreateRentalRequest) (*dto.RentalResponse, error)
	Cancel(ctx context.Context, userID int, rentalID int) error
}
