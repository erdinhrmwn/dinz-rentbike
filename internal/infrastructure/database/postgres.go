package database

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"dinz-rentbike/internal/config"
)

func Connect(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Pass),
		Host:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Path:     cfg.Name,
		RawQuery: "sslmode=disable",
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
