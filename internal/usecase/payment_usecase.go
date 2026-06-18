package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"dinz-rentbike/internal/domain/constants"
	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
	"dinz-rentbike/pkg/utils"
)

type paymentUsecase struct {
	paymentRepo   contract.PaymentRepository
	rentalRepo    contract.RentalRepository
	userRepo      contract.UserRepository
	xenditService contract.XenditService
}

func NewPaymentUsecase(paymentRepo contract.PaymentRepository, rentalRepo contract.RentalRepository, userRepo contract.UserRepository, xenditService contract.XenditService) contract.PaymentUsecase {
	return &paymentUsecase{paymentRepo: paymentRepo, rentalRepo: rentalRepo, userRepo: userRepo, xenditService: xenditService}
}

func (u *paymentUsecase) UserPayments(ctx context.Context, userID int) ([]dto.PaymentResponse, error) {
	payments, err := u.paymentRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var res []dto.PaymentResponse
	for _, p := range payments {
		res = append(res, toPaymentResponse(&p))
	}
	return res, nil
}

func (u *paymentUsecase) PaymentDetail(ctx context.Context, userID int, paymentID int) (*dto.PaymentResponse, error) {
	payment, err := u.paymentRepo.FindByID(ctx, paymentID)
	if err != nil {
		return nil, err
	}

	if payment.UserID != userID {
		return nil, errors.New("payment not found")
	}

	res := toPaymentResponse(payment)
	return &res, nil
}

func (u *paymentUsecase) GetByRentalID(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error) {
	rental, err := u.rentalRepo.FindByID(ctx, rentalID)
	if err != nil {
		return nil, err
	}

	if rental.UserID != userID {
		return nil, errors.New("rental not found")
	}

	payment, err := u.paymentRepo.FindByRentalID(ctx, rentalID)
	if err != nil {
		return nil, err
	}

	res := toPaymentResponse(payment)
	return &res, nil
}

func (u *paymentUsecase) CreatePayment(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error) {
	rental, err := u.rentalRepo.FindByID(ctx, rentalID)
	if err != nil {
		return nil, err
	}

	if rental.UserID != userID {
		return nil, errors.New("rental not found")
	}

	if rental.Status != constants.RentalStatusPending {
		return nil, errors.New("rental is not in pending status")
	}

	payment := &entity.Payment{
		UserID:   userID,
		RentalID: rental.ID,
		Amount:   rental.TotalPrice,
		Status:   constants.PaymentStatusPending,
	}

	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	externalID := fmt.Sprintf("INV-%s-%s", time.Now().Format("20060102"), utils.RandomString(16))
	invoice, err := u.xenditService.CreateInvoice(ctx, &contract.InvoiceRequest{
		ReferenceID: externalID,
		Amount:      rental.TotalPrice,
		Currency:    "IDR",
		PayerEmail:  user.Email,
		Description: fmt.Sprintf("Sewa %s %s", rental.Vehicle.Brand, rental.Vehicle.Name),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}

	payment.XenditInvoiceID = &invoice.InvoiceID
	payment.XenditPaymentURL = &invoice.InvoiceURL

	if err := u.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}

	res := toPaymentResponse(payment)
	return &res, nil
}

func (u *paymentUsecase) CancelPayment(ctx context.Context, userID int, rentalID int) (*dto.PaymentResponse, error) {
	payment, err := u.paymentRepo.FindByRentalID(ctx, rentalID)
	if err != nil {
		return nil, err
	}

	if payment.UserID != userID {
		return nil, errors.New("payment not found")
	}

	if payment.Status != constants.PaymentStatusPending {
		return nil, errors.New("only pending payment can be cancelled")
	}

	if payment.XenditInvoiceID == nil || *payment.XenditInvoiceID == "" {
		return nil, errors.New("no xendit session found")
	}

	if err := u.xenditService.CancelInvoice(ctx, *payment.XenditInvoiceID); err != nil {
		return nil, err
	}

	now := time.Now()
	payment.Status = constants.PaymentStatusFailed
	payment.PaidAt = &now

	if err := u.paymentRepo.Update(ctx, payment); err != nil {
		return nil, err
	}

	res := toPaymentResponse(payment)
	return &res, nil
}

func toPaymentResponse(p *entity.Payment) dto.PaymentResponse {
	res := dto.PaymentResponse{
		ID:               p.ID,
		UserID:           p.UserID,
		RentalID:         p.RentalID,
		Amount:           p.Amount,
		Status:           p.Status,
		PaymentMethod:    p.PaymentMethod,
		XenditInvoiceID:  p.XenditInvoiceID,
		XenditPaymentURL: p.XenditPaymentURL,
		CreatedAt:        p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        p.UpdatedAt.Format(time.RFC3339),
	}

	if p.PaidAt != nil {
		paidAt := p.PaidAt.Format(time.RFC3339)
		res.PaidAt = &paidAt
	}

	return res
}
