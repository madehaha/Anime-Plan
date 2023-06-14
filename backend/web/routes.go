package web

import (
	"backend/web/handler/subject"
	"backend/web/handler/user"
	"backend/web/middleware"
	"github.com/labstack/echo/v4"
)

func AddRouters(app *echo.Echo, middleware middleware.JwtMiddleware, userHandler user.Handler, subjectHandler subject.Handler) {
	app.POST("/register", userHandler.Register)
	app.POST("/login", userHandler.Login)
	app.POST("/cancel", userHandler.Cancel, middleware.UserJWTAuth)
	app.GET("/subject/get", subjectHandler.GetSubject)
	app.GET("/subject/:id", subjectHandler.GetSubjectByID)
	app.POST("/subject/create", subjectHandler.CreateSubject, middleware.UserJWTAuth)

}
