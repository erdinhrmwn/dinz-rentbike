package contract

import (
	"context"

	"dinz-rentbike/internal/domain/dto"
)

type AuthUsecase interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}
