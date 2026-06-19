package bootstrap

import (
	"gorm.io/gorm"

	"dinz-rentbike/internal/config"
	"dinz-rentbike/internal/domain/contract"
	"dinz-rentbike/internal/external/mailjet"
	"dinz-rentbike/internal/external/xendit"
	"dinz-rentbike/internal/infrastructure/database"
	"dinz-rentbike/pkg/logger"
)

type App struct {
	Config        *config.Config
	DB            *gorm.DB
	XenditClient  contract.XenditService
	MailjetClient contract.MailjetService
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

	xenditClient := xendit.NewClient(&cfg.Xendit)
	logger.Log.Info().Msg("xendit client initialized")

	mailjetClient := mailjet.NewClient(&cfg.Mailjet)
	logger.Log.Info().Msg("mailjet client initialized")

	return &App{
		Config:        cfg,
		DB:            db,
		XenditClient:  xenditClient,
		MailjetClient: mailjetClient,
	}
}
