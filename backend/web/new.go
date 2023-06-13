package web

import (
	"backend/internal/config"
	"backend/internal/logger"
	"fmt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	return echo.New()
}

func Start(app *echo.Echo, cfg *config.AppConfig) error {
	addr := cfg.ListenAddr()
	logger.Info(fmt.Sprintf("Server is listening at %s", addr))
	return app.Start(addr)
}
