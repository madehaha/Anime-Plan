package web

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"backend/internal/collection"
	"backend/internal/config"
	"backend/internal/ctrl"
	"backend/internal/driver"
	"backend/internal/subject"
	"backend/internal/user"
	"backend/web"
	"backend/web/util"
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
		fx.Provide(config.AppConfigReader, driver.NewMysqlClient, util.NewJwtUtil),
		fx.Provide(user.NewRepo),
		fx.Provide(subject.NewRepo),
		fx.Provide(collection.NewRepo),

		ctrl.Module,
		web.Module,
		fx.Populate(&cfg, &e),
	).Err()

	if err != nil {
		return err
	}

	return web.Start(e, &cfg)
}
