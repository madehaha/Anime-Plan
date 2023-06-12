package web

import (
	"backend/internal/config"
	"backend/web"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var Command = &cobra.Command{
	Use:   "web",
	Short: "start a web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return start()
	},
}

func start() error {
	// TODO
	var cfg config.AppConfig
	var e *echo.Echo

	err := fx.New(
		fx.NopLogger,

		fx.Provide(
			config.AppConfigReader,
		),

		web.Module,
		fx.Populate(&cfg, &e),
	).Err()

	if err != nil {
		return err
	}

	return web.Start(e, &cfg)
}
