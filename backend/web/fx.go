package web

import (
	"backend/web/handler"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"web",
	fx.Provide(New),
	handler.Module,
	fx.Invoke(AddRouters),
)
