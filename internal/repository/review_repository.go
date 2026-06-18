package repository

import (
	"context"

	"gorm.io/gorm"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/entity"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) contract.ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) FindByUserID(ctx context.Context, userID int) ([]entity.Review, error) {
	var reviews []entity.Review
	query := r.db.WithContext(ctx).
		Preload("Vehicle").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&reviews)
	if query.Error != nil {
		return nil, query.Error
	}
	return reviews, nil
}

func (r *reviewRepository) FindByRentalID(ctx context.Context, rentalID int) (*entity.Review, error) {
	var review entity.Review
	query := r.db.WithContext(ctx).Where("rental_id = ?", rentalID).First(&review)
	if query.Error != nil {
		return nil, query.Error
	}
	return &review, nil
}

func (r *reviewRepository) FindByVehicleID(ctx context.Context, vehicleID int) ([]entity.Review, error) {
	var reviews []entity.Review
	query := r.db.WithContext(ctx).
		Preload("User").
		Where("vehicle_id = ?", vehicleID).
		Order("created_at DESC").
		Find(&reviews)
	if query.Error != nil {
		return nil, query.Error
	}
	return reviews, nil
}

func (r *reviewRepository) Create(ctx context.Context, review *entity.Review) error {
	query := r.db.WithContext(ctx).Create(review)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *reviewRepository) Update(ctx context.Context, review *entity.Review) error {
	query := r.db.WithContext(ctx).Save(review)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *reviewRepository) Delete(ctx context.Context, id int) error {
	query := r.db.WithContext(ctx).Delete(&entity.Review{}, id)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
