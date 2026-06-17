package bootstrap

import (
	"dinz-rentbike/internal/delivery/http/middleware"
	"dinz-rentbike/pkg/logger"
	"dinz-rentbike/pkg/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func (a *App) Run() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(echoMiddleware.RequestID())
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.RequestLoggerMiddleware())

	e.GET("/", func(c echo.Context) error {
		return response.SuccessResponse(c, http.StatusOK, "Hello World!", nil)
	})

	addr := fmt.Sprintf("%s:%d", a.Config.App.Host, a.Config.App.Port)
	logger.Log.Info().Str("address", addr).Msg("starting server")
	if err := e.Start(addr); err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to start server")
	}
}
