package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDB),
	fx.Provide(NewEnv),
	fx.Provide(NewAMQP),
)
