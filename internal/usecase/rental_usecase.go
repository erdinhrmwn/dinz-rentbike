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

func (u *rentalUsecase) RentalDetail(ctx context.Context, userID int, rentalID int) (*dto.RentalResponse, error) {
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

func (u *rentalUsecase) UserRentals(ctx context.Context, userID int) ([]dto.RentalResponse, error) {
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

func (u *rentalUsecase) CreateRental(ctx context.Context, userID int, req *dto.CreateRentalRequest) (*dto.RentalResponse, error) {
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

func (u *rentalUsecase) CancelRental(ctx context.Context, userID int, rentalID int) error {
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
	return u.vehicleRepo.Update(ctx, vehicle)
}

func (u *rentalUsecase) ListAll(ctx context.Context) ([]dto.RentalResponse, error) {
	rentals, err := u.rentalRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.RentalResponse
	for _, r := range rentals {
		res = append(res, toRentalResponse(&r))
	}
	return res, nil
}

func (u *rentalUsecase) AdminDetail(ctx context.Context, rentalID int) (*dto.RentalResponse, error) {
	rental, err := u.rentalRepo.FindByID(ctx, rentalID)
	if err != nil {
		return nil, err
	}

	res := toRentalResponse(rental)
	return &res, nil
}

func (u *rentalUsecase) UpdateStatus(ctx context.Context, rentalID int, status string) (*dto.RentalResponse, error) {
	rental, err := u.rentalRepo.FindByID(ctx, rentalID)
	if err != nil {
		return nil, err
	}

	rental.Status = status
	if err := u.rentalRepo.Update(ctx, rental); err != nil {
		return nil, err
	}

	if status == constants.RentalStatusCancelled || status == constants.RentalStatusCompleted {
		vehicle, err := u.vehicleRepo.FindByID(ctx, rental.VehicleID)
		if err == nil {
			vehicle.Status = constants.VehicleStatusAvailable
			u.vehicleRepo.Update(ctx, vehicle)
		}
	}

	res := toRentalResponse(rental)
	return &res, nil
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
		res.Vehicle = dto.VehicleResponse{
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

	if r.Payment != nil {
		p := r.Payment
		payment := dto.PaymentResponse{
			ID:               p.ID,
			UserID:           p.UserID,
			RentalID:         p.RentalID,
			Amount:           p.Amount,
			Status:           p.Status,
			XenditInvoiceID:  p.XenditInvoiceID,
			XenditPaymentURL: p.XenditPaymentURL,
			CreatedAt:        p.CreatedAt.Format(time.RFC3339),
			UpdatedAt:        p.UpdatedAt.Format(time.RFC3339),
		}
		if p.PaidAt != nil {
			paidAt := p.PaidAt.Format(time.RFC3339)
			payment.PaidAt = &paidAt
		}
		res.Payment = &payment
	}

	if r.Review != nil {
		rv := r.Review
		review := dto.ReviewResponse{
			ID:        rv.ID,
			UserID:    rv.UserID,
			VehicleID: rv.VehicleID,
			RentalID:  rv.RentalID,
			Rating:    rv.Rating,
			Comment:   rv.Comment,
			CreatedAt: rv.CreatedAt.Format(time.RFC3339),
			UpdatedAt: rv.UpdatedAt.Format(time.RFC3339),
		}
		res.Review = &review
	}

	return res
}
