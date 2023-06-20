package web

import (
	"net/http"

	"backend/web/handler/collection"
	"backend/web/handler/subject"
	"backend/web/handler/user"
	"backend/web/middleware"
	"backend/web/util"

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
	app.GET("/user/:member_id", userHandler.GetMember)

	// Subject
	app.GET("/subject/get", subjectHandler.GetSubject)
	app.POST(
		"/subject/create", subjectHandler.CreateSubject,
		middleware.WikiJWTAuth,
	) // verify if having permission to create subject

	app.GET("/subject/:subject_id", subjectHandler.GetSubjectByID)

	// TODO update subject_field
	app.POST("/collection/:subject_id", collectionHandler.AddCollection, middleware.UserJWTAuth)
	app.PATCH("/collection/:subject_id", collectionHandler.UpdateCollection, middleware.UserJWTAuth)
	app.DELETE("/collection/:subject_id", collectionHandler.DeleteCollection, middleware.UserJWTAuth)
	// TODO
	app.GET("/:subject_id/subject/comment", collectionHandler.GetCommentsBySubjectID)
	app.GET("/:member_id/member/comment", collectionHandler.GetMemberComment)
	app.Any(
		"/*", func(c echo.Context) error {
			return util.Error(c, http.StatusNotFound, "")
		},
	)

}
