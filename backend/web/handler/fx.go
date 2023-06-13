package handler

import (
	"backend/web/handler/user"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"handler",
	fx.Provide(user.NewUserHandler),
)
