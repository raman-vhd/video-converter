package route

import (
	"github.com/raman-vhd/video-converter/internal/api/controller"
	"github.com/raman-vhd/video-converter/internal/lib"
)

type templateRoute struct {
    handler lib.RequestHandler
    ctrl controller.ITemplateController
}

func NewTemplate(
    handler lib.RequestHandler,
    ctrl controller.ITemplateController,
) templateRoute{
    return templateRoute{
        handler: handler,
        ctrl: ctrl,
    }
}

func (a templateRoute) Setup() {
    // api := a.handler.Echo.Group("/api")
    
    // api.GET("/template", a.ctrl.Action)
}
