package web

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"backend/internal/config"
	"backend/internal/logger"
)

func New() *echo.Echo {
	return echo.New()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Start(app *echo.Echo, cfg *config.AppConfig) error {
	addr := cfg.ListenAddr()
	logger.Info(fmt.Sprintf("Server is listening at %s", addr))
	app.Validator = &CustomValidator{validator: validator.New()}
	app.Static("public", cfg.StaticDirectory)
	return app.Start(addr)
}
