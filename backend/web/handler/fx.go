package handler

import (
	"go.uber.org/fx"

	"backend/web/handler/collection"
	"backend/web/handler/subject"
	"backend/web/handler/user"
)

var Module = fx.Module(
	"handler",
	fx.Provide(user.NewUserHandler),
	fx.Provide(subject.NewSubjectHandler),
	fx.Provide(collection.NewCollectionHandler),
)
