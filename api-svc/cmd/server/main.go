package main

import (
	"context"
	"log"

	"github.com/raman-vhd/video-converter/internal/api/controller"
	"github.com/raman-vhd/video-converter/internal/repository"
	"github.com/raman-vhd/video-converter/internal/api/route"
	"github.com/raman-vhd/video-converter/internal/service"
	"github.com/raman-vhd/video-converter/internal/lib"
	"go.uber.org/fx"
)

var Module = fx.Options(
    repository.Module,
    controller.Module,
    service.Module,
    route.Module,
    lib.Module,
    )

func main() {
    app := fx.New(
        Module,
        fx.Invoke(bootstrap),
        )
    ctx := context.Background()
    err := app.Start(ctx)
    defer app.Stop(ctx)
    if err != nil {
        log.Fatalf("failed starting app: %v", err)
    }
}

func bootstrap(
    routes route.Routes,
    env lib.Env,
    router lib.RequestHandler,
) {
    routes.Setup()
    
    log.Println("Running Server")
    
    router.Echo.Start(":" + env.ServerPort)
}
