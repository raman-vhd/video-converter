package route

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewTemplate,
		NewVideo,
	),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	template templateRoute,
	video videoRoute,
) Routes {
	return Routes{
		template,
		video,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
