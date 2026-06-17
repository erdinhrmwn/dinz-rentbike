package bootstrap

import (
	"context"
	"dinz-rentbike/internal/delivery/http/middleware"
	"dinz-rentbike/pkg/logger"
	"dinz-rentbike/pkg/response"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func (a *App) Run() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())
	e.Use(middleware.RequestLoggerMiddleware())

	e.GET("/", func(c echo.Context) error {
		return response.SuccessResponse(c, http.StatusOK, "Hello World!", nil)
	})

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
