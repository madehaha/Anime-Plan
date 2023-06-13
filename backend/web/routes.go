package web

import (
	"backend/web/handler/user"
	"github.com/labstack/echo/v4"
)

func AddRouters(app *echo.Echo, userHandler user.Handler) {
	userWithoutJWT := app.Group("user")
	//user := app.Group("user")
	userWithoutJWT.POST("/register", userHandler.Register)
	//user.POST("/subject/create",userHandler.CreateSubject)
}
