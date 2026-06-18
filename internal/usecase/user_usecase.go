package usecase

import (
	"context"
	"errors"
	"time"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/pkg/utils"
)

type userUsecase struct {
	userRepo contract.UserRepository
}

func NewUserUsecase(userRepo contract.UserRepository) contract.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetProfile(ctx context.Context, userID int) (*dto.UserResponse, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, userID int, req *dto.UpdateProfileRequest) error {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	user.Name = req.Name
	user.Phone = req.Phone

	if err := u.userRepo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) ChangePassword(ctx context.Context, userID int, req *dto.ChangePasswordRequest) error {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if req.NewPassword != req.ConfirmPassword {
		return errors.New("new password and confirm password do not match")
	}

	if !utils.ValidatePassword(req.OldPassword, user.Password) {
		return errors.New("old password is incorrect")
	}

	hashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashed
	if err := u.userRepo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
