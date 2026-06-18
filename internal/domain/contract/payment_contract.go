package contract

import (
	"context"

	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type PaymentRepository interface {
	FindByID(ctx context.Context, id int) (*entity.Payment, error)
	FindByUserID(ctx context.Context, userID int) ([]entity.Payment, error)
	FindByRentalID(ctx context.Context, rentalID int) (*entity.Payment, error)
	Create(ctx context.Context, payment *entity.Payment) error
	Update(ctx context.Context, payment *entity.Payment) error
}

type PaymentUsecase interface {
	GetByUserID(ctx context.Context, userID int) ([]dto.PaymentResponse, error)
	GetByID(ctx context.Context, userID int, paymentID int) (*dto.PaymentResponse, error)
	GetByRentalID(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error)
	Create(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error)
	Cancel(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error)
}
