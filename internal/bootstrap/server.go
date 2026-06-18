package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"dinz-rentbike/internal/delivery/http/handler"
	"dinz-rentbike/internal/delivery/http/middleware"
	"dinz-rentbike/internal/repository"
	"dinz-rentbike/internal/usecase"
	"dinz-rentbike/pkg/jwt"
	"dinz-rentbike/pkg/logger"
	"dinz-rentbike/pkg/response"
)

func (a *App) Run() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())
	e.Use(middleware.RequestLoggerMiddleware())

	// Auth
	authManager := jwt.NewAuthManager(&a.Config.Jwt)
	authMiddleware := middleware.AuthMiddleware(authManager)

	// Repositories
	userRepo := repository.NewUserRepository(a.DB)
	vehicleRepo := repository.NewVehicleRepository(a.DB)
	rentalRepo := repository.NewRentalRepository(a.DB)
	paymentRepo := repository.NewPaymentRepository(a.DB)
	reviewRepo := repository.NewReviewRepository(a.DB)

	// Usecases
	authUsecase := usecase.NewAuthUsecase(userRepo, authManager)
	userUsecase := usecase.NewUserUsecase(userRepo)
	vehicleUsecase := usecase.NewVehicleUsecase(vehicleRepo)
	rentalUsecase := usecase.NewRentalUsecase(rentalRepo, vehicleRepo)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, rentalRepo, userRepo, a.XenditClient)
	reviewUsecase := usecase.NewReviewUsecase(reviewRepo, rentalRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	vehicleHandler := handler.NewVehicleHandler(vehicleUsecase)
	rentalHandler := handler.NewRentalHandler(rentalUsecase)
	paymentHandler := handler.NewPaymentHandler(paymentUsecase)
	reviewHandler := handler.NewReviewHandler(reviewUsecase)

	// Routes
	e.GET("/", func(c echo.Context) error {
		return response.SuccessResponse(c, http.StatusOK, "Hello World!", nil)
	})

	api := e.Group("/api/v1")

	authGroup := api.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	userGroup := api.Group("/users", authMiddleware)
	userHandler.RegisterRoutes(userGroup)

	vehicleGroup := api.Group("/vehicles", authMiddleware)
	vehicleHandler.RegisterRoutes(vehicleGroup)

	rentalGroup := api.Group("/rentals", authMiddleware)
	rentalHandler.RegisterRoutes(rentalGroup)

	paymentGroup := api.Group("/payments", authMiddleware)
	paymentHandler.RegisterRoutes(paymentGroup)

	reviewGroup := api.Group("/reviews", authMiddleware)
	reviewHandler.RegisterRoutes(reviewGroup)

	addr := fmt.Sprintf("%s:%d", a.Config.App.Host, a.Config.App.Port)

	// Start server in goroutine
	go func() {
		logger.Log.Info().Str("address", addr).Msg("starting server")
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info().Msg("shutting down server")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Log.Fatal().Err(err).Msg("server forced to shutdown")
	}

	logger.Log.Info().Msg("server stopped gracefully")
}
