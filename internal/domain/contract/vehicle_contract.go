package contract

import (
	"context"

	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type VehicleRepository interface {
	FindAll(ctx context.Context) ([]entity.Vehicle, error)
	FindByID(ctx context.Context, id int) (*entity.Vehicle, error)
	Update(ctx context.Context, vehicle *entity.Vehicle) error
}

type VehicleUsecase interface {
	GetAll(ctx context.Context) ([]dto.VehicleResponse, error)
	GetByID(ctx context.Context, id int) (*dto.VehicleResponse, error)
}
