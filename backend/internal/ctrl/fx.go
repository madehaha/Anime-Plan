package ctrl

import "go.uber.org/fx"

var Module = fx.Module(
	"ctrl",
	fx.Provide(NewUserCtrl),
)
