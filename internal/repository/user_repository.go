package repository

import (
	"context"

	"gorm.io/gorm"

	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/domain/entity"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) contract.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	query := r.db.WithContext(ctx).First(&user, id)
	if query.Error != nil {
		return nil, query.Error
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	query := r.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if query.Error != nil {
		return nil, query.Error
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := r.db.WithContext(ctx).Create(user)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	query := r.db.WithContext(ctx).Save(user)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
