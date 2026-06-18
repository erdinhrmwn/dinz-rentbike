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
	FindAll(ctx context.Context) ([]entity.Payment, error)
	Create(ctx context.Context, payment *entity.Payment) error
	Update(ctx context.Context, payment *entity.Payment) error
}

type PaymentUsecase interface {
	UserPayments(ctx context.Context, userID int) ([]dto.PaymentResponse, error)
	PaymentDetail(ctx context.Context, userID int, paymentID int) (*dto.PaymentResponse, error)
	GetByRentalID(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error)
	CreatePayment(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error)
	CancelPayment(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error)
	ListAll(ctx context.Context) ([]dto.PaymentResponse, error)
	AdminDetail(ctx context.Context, paymentID int) (*dto.PaymentResponse, error)
}
