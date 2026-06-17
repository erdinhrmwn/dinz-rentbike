package contract

import (
	"context"

	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
}

type UserUsecase interface {
	GetProfile(ctx context.Context, userID int) (*dto.ProfileResponse, error)
	UpdateProfile(ctx context.Context, userID int, req *dto.UpdateProfileRequest) error
	ChangePassword(ctx context.Context, userID int, req *dto.ChangePasswordRequest) error
}
