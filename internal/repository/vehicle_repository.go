package repository

import (
	"context"

	"gorm.io/gorm"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/entity"
)

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) contract.VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) FindAll(ctx context.Context) ([]entity.Vehicle, error) {
	var vehicles []entity.Vehicle
	query := r.db.WithContext(ctx).Find(&vehicles)
	if query.Error != nil {
		return nil, query.Error
	}
	return vehicles, nil
}

func (r *vehicleRepository) FindByID(ctx context.Context, id int) (*entity.Vehicle, error) {
	var vehicle entity.Vehicle
	query := r.db.WithContext(ctx).First(&vehicle, id)
	if query.Error != nil {
		return nil, query.Error
	}
	return &vehicle, nil
}

func (r *vehicleRepository) Update(ctx context.Context, vehicle *entity.Vehicle) error {
	query := r.db.WithContext(ctx).Save(vehicle)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
