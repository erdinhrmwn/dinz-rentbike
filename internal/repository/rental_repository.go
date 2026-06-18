package repository

import (
	"context"

	"gorm.io/gorm"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/entity"
)

type rentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) contract.RentalRepository {
	return &rentalRepository{db: db}
}

func (r *rentalRepository) FindByID(ctx context.Context, id int) (*entity.Rental, error) {
	var rental entity.Rental
	query := r.db.WithContext(ctx).Preload("Vehicle").Preload("Payment").Preload("Review").First(&rental, id)
	if query.Error != nil {
		return nil, query.Error
	}
	return &rental, nil
}

func (r *rentalRepository) FindByUserID(ctx context.Context, userID int) ([]entity.Rental, error) {
	var rentals []entity.Rental
	query := r.db.WithContext(ctx).
		Preload("Vehicle").
		Preload("Payment").
		Preload("Review").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&rentals)
	if query.Error != nil {
		return nil, query.Error
	}
	return rentals, nil
}

func (r *rentalRepository) Create(ctx context.Context, rental *entity.Rental) error {
	query := r.db.WithContext(ctx).Create(rental)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *rentalRepository) Update(ctx context.Context, rental *entity.Rental) error {
	query := r.db.WithContext(ctx).Save(rental)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
