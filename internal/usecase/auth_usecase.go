package usecase

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"dinz-rentbike/internal/domain/constants"
	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
	"dinz-rentbike/pkg/utils"
)

type authUsecase struct {
	userRepo contract.UserRepository
}

func NewAuthUsecase(userRepo contract.UserRepository) contract.AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

func (u *authUsecase) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashed,
		Role:     constants.UserRoleCustomer,
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		User: dto.ProfileResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
			Role:  user.Role,
		},
	}, nil
}

func (u *authUsecase) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !utils.ValidatePassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	return &dto.LoginResponse{
		Token: "jwt-token-placeholder",
	}, nil
}
