package repository

import (
	"context"

	"gorm.io/gorm"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/entity"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) contract.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) FindByID(ctx context.Context, id int) (*entity.Payment, error) {
	var payment entity.Payment
	query := r.db.WithContext(ctx).First(&payment, id)
	if query.Error != nil {
		return nil, query.Error
	}
	return &payment, nil
}

func (r *paymentRepository) FindByUserID(ctx context.Context, userID int) ([]entity.Payment, error) {
	var payments []entity.Payment
	query := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&payments)
	if query.Error != nil {
		return nil, query.Error
	}
	return payments, nil
}

func (r *paymentRepository) FindByRentalID(ctx context.Context, rentalID int) (*entity.Payment, error) {
	var payment entity.Payment
	query := r.db.WithContext(ctx).Where("rental_id = ?", rentalID).First(&payment)
	if query.Error != nil {
		return nil, query.Error
	}
	return &payment, nil
}

func (r *paymentRepository) Create(ctx context.Context, payment *entity.Payment) error {
	query := r.db.WithContext(ctx).Create(payment)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *paymentRepository) Update(ctx context.Context, payment *entity.Payment) error {
	query := r.db.WithContext(ctx).Save(payment)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
