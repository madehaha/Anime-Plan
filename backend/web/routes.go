package web

import (
	"backend/web/handler/user"
	"github.com/labstack/echo/v4"
)

func AddRouters(app *echo.Echo, userHandler user.Handler) {
	app.POST("/register", userHandler.Register)
	app.POST("/login", userHandler.Login)
}
