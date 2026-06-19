package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"dinz-rentbike/internal/config"
	"dinz-rentbike/internal/domain/constants"
	"dinz-rentbike/internal/domain/dto"
	"dinz-rentbike/internal/domain/entity"
	"dinz-rentbike/internal/usecase"
	"dinz-rentbike/pkg/jwt"
	"dinz-rentbike/pkg/utils"
	"dinz-rentbike/tests/mocks"
)

func TestRegister_Success(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().Create(mock.Anything, mock.Anything).Return(nil)

	authManager := jwt.NewAuthManager(&config.JwtConfig{Secret: "test-secret"})
	authUsecase := usecase.NewAuthUsecase(repo, authManager, nil)

	req := &dto.RegisterRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Phone:    "081234567890",
		Password: "password123",
	}

	res, err := authUsecase.Register(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, req.Name, res.User.Name)
	assert.Equal(t, req.Email, res.User.Email)
	assert.Equal(t, constants.UserRoleCustomer, res.User.Role)
}

func TestRegister_DuplicateEmail(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().Create(mock.Anything, mock.Anything).Return(nil).Once()
	repo.EXPECT().Create(mock.Anything, mock.Anything).Return(errors.New("duplicate key")).Once()

	authManager := jwt.NewAuthManager(&config.JwtConfig{Secret: "test-secret"})
	authUsecase := usecase.NewAuthUsecase(repo, authManager, nil)

	req := &dto.RegisterRequest{
		Name:     "Test User",
		Email:    "dupe@example.com",
		Phone:    "081234567890",
		Password: "password123",
	}

	_, err := authUsecase.Register(context.Background(), req)
	require.NoError(t, err)

	_, err = authUsecase.Register(context.Background(), req)
	assert.Error(t, err)
}

func TestLogin_Success(t *testing.T) {
	hashed, _ := utils.HashPassword("password123")

	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().FindByEmail(mock.Anything, "login@example.com").
		Return(&entity.User{
			ID:       1,
			Name:     "Login User",
			Email:    "login@example.com",
			Password: string(hashed),
			Role:     constants.UserRoleCustomer,
		}, nil)

	authManager := jwt.NewAuthManager(&config.JwtConfig{Secret: "test-secret"})
	authUsecase := usecase.NewAuthUsecase(repo, authManager, nil)

	req := &dto.LoginRequest{
		Email:    "login@example.com",
		Password: "password123",
	}

	res, err := authUsecase.Login(context.Background(), req)
	require.NoError(t, err)
	assert.NotEmpty(t, res.Token)
}

func TestLogin_InvalidEmail(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().
		FindByEmail(mock.Anything, "unknown@example.com").
		Return(nil, errors.New("record not found"))

	authManager := jwt.NewAuthManager(&config.JwtConfig{Secret: "test-secret"})
	authUsecase := usecase.NewAuthUsecase(repo, authManager, nil)

	req := &dto.LoginRequest{
		Email:    "unknown@example.com",
		Password: "password123",
	}

	_, err := authUsecase.Login(context.Background(), req)
	assert.Error(t, err)
}

func TestLogin_InvalidPassword(t *testing.T) {
	hashed, _ := utils.HashPassword("password123")

	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().FindByEmail(mock.Anything, "user@example.com").
		Return(&entity.User{
			ID:       1,
			Name:     "User",
			Email:    "user@example.com",
			Password: string(hashed),
			Role:     constants.UserRoleCustomer,
		}, nil)

	authManager := jwt.NewAuthManager(&config.JwtConfig{Secret: "test-secret"})
	authUsecase := usecase.NewAuthUsecase(repo, authManager, nil)

	req := &dto.LoginRequest{
		Email:    "user@example.com",
		Password: "wrongpassword",
	}

	_, err := authUsecase.Login(context.Background(), req)
	assert.Error(t, err)
}
