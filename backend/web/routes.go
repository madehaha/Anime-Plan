package web

import (
	"github.com/labstack/echo/v4"

	"backend/web/handler/user"
	"backend/web/middleware"
)

func AddRouters(
	app *echo.Echo, middleware middleware.JwtMiddleware,
	userHandler user.Handler,
) {
	app.POST("/register", userHandler.Register)
	app.POST("/login", userHandler.Login)
	app.POST("/cancel", userHandler.Cancel, middleware.UserJWTAuth)
	app.GET("/user/:id/avatar", userHandler.GetAvatar)
	app.GET("/me", userHandler.GetMe, middleware.UserJWTAuth)
}
