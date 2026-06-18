package usecase

import (
	"context"
	"errors"
	"math"
	"time"

	"dinz-rentbike/internal/domain/constants"
	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type rentalUsecase struct {
	rentalRepo  contract.RentalRepository
	vehicleRepo contract.VehicleRepository
}

func NewRentalUsecase(rentalRepo contract.RentalRepository, vehicleRepo contract.VehicleRepository) contract.RentalUsecase {
	return &rentalUsecase{rentalRepo: rentalRepo, vehicleRepo: vehicleRepo}
}

func (u *rentalUsecase) GetByID(ctx context.Context, userID int, rentalID int) (*dto.RentalResponse, error) {
	rental, err := u.rentalRepo.FindByID(ctx, rentalID)
	if err != nil {
		return nil, err
	}

	if rental.UserID != userID {
		return nil, errors.New("rental not found")
	}

	res := toRentalResponse(rental)
	return &res, nil
}

func (u *rentalUsecase) GetByUserID(ctx context.Context, userID int) ([]dto.RentalResponse, error) {
	rentals, err := u.rentalRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var res []dto.RentalResponse
	for _, r := range rentals {
		res = append(res, toRentalResponse(&r))
	}
	return res, nil
}

func (u *rentalUsecase) Create(ctx context.Context, userID int, req *dto.CreateRentalRequest) (*dto.RentalResponse, error) {
	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		return nil, errors.New("invalid start time format")
	}

	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		return nil, errors.New("invalid end time format")
	}

	if endTime.Before(startTime) || endTime.Equal(startTime) {
		return nil, errors.New("end time must be after start time")
	}

	vehicle, err := u.vehicleRepo.FindByID(ctx, req.VehicleID)
	if err != nil {
		return nil, err
	}

	if vehicle.Status != constants.VehicleStatusAvailable {
		return nil, errors.New("vehicle is not available")
	}

	duration := endTime.Sub(startTime)
	totalHours := int(math.Ceil(duration.Hours()))
	totalPrice := float64(totalHours) * vehicle.PricePerHour

	rental := &entity.Rental{
		UserID:     userID,
		VehicleID:  vehicle.ID,
		StartTime:  startTime,
		EndTime:    endTime,
		TotalHours: totalHours,
		TotalPrice: totalPrice,
		Status:     constants.RentalStatusPending,
	}

	if err := u.rentalRepo.Create(ctx, rental); err != nil {
		return nil, err
	}

	vehicle.Status = constants.VehicleStatusRented
	if err := u.vehicleRepo.Update(ctx, vehicle); err != nil {
		return nil, err
	}

	rental.Vehicle = vehicle
	res := toRentalResponse(rental)
	return &res, nil
}

func (u *rentalUsecase) Cancel(ctx context.Context, userID int, rentalID int) error {
	rental, err := u.rentalRepo.FindByID(ctx, rentalID)
	if err != nil {
		return err
	}

	if rental.UserID != userID {
		return errors.New("rental not found")
	}

	if rental.Status != constants.RentalStatusPending && rental.Status != constants.RentalStatusActive {
		return errors.New("rental cannot be cancelled")
	}

	rental.Status = constants.RentalStatusCancelled
	if err := u.rentalRepo.Update(ctx, rental); err != nil {
		return err
	}

	vehicle, err := u.vehicleRepo.FindByID(ctx, rental.VehicleID)
	if err != nil {
		return err
	}

	vehicle.Status = constants.VehicleStatusAvailable
	if err := u.vehicleRepo.Update(ctx, vehicle); err != nil {
		return err
	}

	return nil
}

func toRentalResponse(r *entity.Rental) dto.RentalResponse {
	res := dto.RentalResponse{
		ID:         r.ID,
		UserID:     r.UserID,
		VehicleID:  r.VehicleID,
		StartTime:  r.StartTime.Format(time.RFC3339),
		EndTime:    r.EndTime.Format(time.RFC3339),
		TotalHours: r.TotalHours,
		TotalPrice: r.TotalPrice,
		Status:     r.Status,
		CreatedAt:  r.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  r.UpdatedAt.Format(time.RFC3339),
	}

	if r.Vehicle != nil {
		res.Vehicle = &dto.VehicleResponse{
			ID:           r.Vehicle.ID,
			Type:         r.Vehicle.Type,
			Brand:        r.Vehicle.Brand,
			Name:         r.Vehicle.Name,
			Category:     r.Vehicle.Category,
			Description:  r.Vehicle.Description,
			ImageURL:     r.Vehicle.ImageURL,
			PricePerHour: r.Vehicle.PricePerHour,
			Status:       r.Vehicle.Status,
			CreatedAt:    r.Vehicle.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    r.Vehicle.UpdatedAt.Format(time.RFC3339),
		}
	}

	return res
}
