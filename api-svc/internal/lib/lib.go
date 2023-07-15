package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewDB,
		NewEnv,
		NewRequestHandler,
		NewAMQP,
	),
)
