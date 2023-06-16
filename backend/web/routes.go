package web

import (
	"backend/web/handler/subject"
	"github.com/labstack/echo/v4"

	"backend/web/handler/user"
	"backend/web/middleware"
)

func AddRouters(app *echo.Echo, middleware middleware.JwtMiddleware, userHandler user.Handler, subjectHandler subject.Handler) {
	app.POST("/register", userHandler.Register)
	app.POST("/login", userHandler.Login)
	app.POST("/cancel", userHandler.Cancel, middleware.UserJWTAuth)
	app.PUT("/modify", userHandler.ModifyInfo, middleware.UserJWTAuth)
	app.GET("/me", userHandler.GetMe, middleware.UserJWTAuth)
	app.GET("/user/:id/avatar", userHandler.GetAvatar)

	app.GET("/subject/get", subjectHandler.GetSubject)
	app.GET("/subject/:id", subjectHandler.GetSubjectByID)
	app.POST("/subject/create", subjectHandler.CreateSubject, middleware.UserJWTAuth)
	app.GET("/subject/search", subjectHandler.SearchSubject)

	app.POST("/collection/:id", subjectHandler.AddOrUpdateCollection, middleware.UserJWTAuth)

}
