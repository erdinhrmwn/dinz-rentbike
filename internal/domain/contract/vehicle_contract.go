package contract

import (
	"context"

	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type VehicleRepository interface {
	FindAll(ctx context.Context) ([]entity.Vehicle, error)
	FindByID(ctx context.Context, id int) (*entity.Vehicle, error)
	Create(ctx context.Context, vehicle *entity.Vehicle) error
	Update(ctx context.Context, vehicle *entity.Vehicle) error
	Delete(ctx context.Context, id int) error
}

type VehicleUsecase interface {
	VehicleList(ctx context.Context) ([]dto.VehicleResponse, error)
	VehicleDetail(ctx context.Context, id int) (*dto.VehicleResponse, error)
	Create(ctx context.Context, req *dto.CreateVehicleRequest) (*dto.VehicleResponse, error)
	Update(ctx context.Context, id int, req *dto.UpdateVehicleRequest) (*dto.VehicleResponse, error)
	Delete(ctx context.Context, id int) error
}
