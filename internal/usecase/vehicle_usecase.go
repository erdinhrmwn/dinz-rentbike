package usecase

import (
	"context"
	"time"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type vehicleUsecase struct {
	vehicleRepo contract.VehicleRepository
}

func NewVehicleUsecase(vehicleRepo contract.VehicleRepository) contract.VehicleUsecase {
	return &vehicleUsecase{vehicleRepo: vehicleRepo}
}

func (u *vehicleUsecase) GetAll(ctx context.Context) ([]dto.VehicleResponse, error) {
	vehicles, err := u.vehicleRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.VehicleResponse
	for _, v := range vehicles {
		res = append(res, toVehicleResponse(&v))
	}
	return res, nil
}

func (u *vehicleUsecase) GetByID(ctx context.Context, id int) (*dto.VehicleResponse, error) {
	vehicle, err := u.vehicleRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := toVehicleResponse(vehicle)
	return &res, nil
}

func toVehicleResponse(v *entity.Vehicle) dto.VehicleResponse {
	res := dto.VehicleResponse{
		ID:           v.ID,
		Type:         v.Type,
		Brand:        v.Brand,
		Name:         v.Name,
		Category:     v.Category,
		Description:  v.Description,
		ImageURL:     v.ImageURL,
		PricePerHour: v.PricePerHour,
		Status:       v.Status,
		CreatedAt:    v.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    v.UpdatedAt.Format(time.RFC3339),
	}

	for _, r := range v.Reviews {
		res.Reviews = append(res.Reviews, dto.ReviewResponse{
			ID:        r.ID,
			UserID:    r.UserID,
			VehicleID: r.VehicleID,
			RentalID:  r.RentalID,
			Rating:    r.Rating,
			Comment:   r.Comment,
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
			UpdatedAt: r.UpdatedAt.Format(time.RFC3339),
		})
	}

	return res
}
