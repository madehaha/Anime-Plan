package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddRouters(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error {
		err := c.JSON(http.StatusOK, string("Hello"))
		if err != nil {
			return err
		}
		return nil
	})
}
