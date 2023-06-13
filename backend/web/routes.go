package web

import (
	"backend/web/handler/user"
	"backend/web/middleware"
	"github.com/labstack/echo/v4"
)

func AddRouters(app *echo.Echo, middleware middleware.JwtMiddleware, userHandler user.Handler) {
	app.POST("/register", userHandler.Register)
	app.POST("/login", userHandler.Login)
	app.POST("/cancel", userHandler.Cancel, middleware.UserJWTAuth)
}
