package web

import (
	"backend/web/handler"
	"backend/web/middleware"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"web",
	fx.Provide(New),
	fx.Provide(middleware.NewJwtMiddleware),
	handler.Module,
	fx.Invoke(AddRouters),
)
