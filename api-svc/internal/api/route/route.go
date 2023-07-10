package route

import "go.uber.org/fx"

var Module = fx.Options(
    fx.Provide(
        NewTemplate,
        ),
    fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
    Setup()
}

func NewRoutes(
    template templateRoute ,
) Routes {
    return Routes{
        template,
    }
}

func (r Routes) Setup() {
    for _, route := range r {
        route.Setup()
    }
}
