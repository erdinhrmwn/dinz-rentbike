package bootstrap

import (
	"dinz-rentbike/internal/config"
	"dinz-rentbike/internal/infrastructure/database"
	"dinz-rentbike/pkg/logger"

	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
}

func Init() *App {
	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to load config")
		return nil
	}
	logger.Log.Info().Msg("config loaded")

	db, err := database.Connect(&cfg.Database)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to connect to database")
		return nil
	}
	logger.Log.Info().Msg("database connected")

	return &App{
		Config: cfg,
		DB:     db,
	}
}
