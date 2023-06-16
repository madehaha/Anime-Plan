package web

import (
	"backend/web/handler/collection"
	"backend/web/handler/subject"
	"backend/web/handler/user"
	"backend/web/middleware"

	"github.com/labstack/echo/v4"
)

func AddRouters(
	app *echo.Echo, middleware middleware.JwtMiddleware, userHandler user.Handler, subjectHandler subject.Handler,
	collectionHandler collection.Handler,
) {
	// User
	app.POST("/register", userHandler.Register)
	app.POST("/login", userHandler.Login)
	app.POST("/cancel", userHandler.Cancel, middleware.UserJWTAuth)
	app.PUT("/modify", userHandler.ModifyInfo, middleware.UserJWTAuth)
	app.GET("/me", userHandler.GetMe, middleware.UserJWTAuth)
	app.GET("/user/:id/avatar", userHandler.GetAvatar)

	// Subject
	app.GET("/subject/get", subjectHandler.GetSubject)
	app.GET("/subject/:subject_id", subjectHandler.GetSubjectByID)
	app.POST("/subject/create", subjectHandler.CreateSubject, middleware.UserJWTAuth)
	app.GET("/subject/search", subjectHandler.SearchSubject)

	app.POST("/collection/:subject_id", collectionHandler.AddCollection, middleware.UserJWTAuth)
	app.PATCH("/collection/:subject_id", collectionHandler.UpdateCollection, middleware.UserJWTAuth)
}
