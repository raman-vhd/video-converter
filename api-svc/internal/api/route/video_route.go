package route

import (
	"github.com/raman-vhd/video-converter/internal/api/controller"
	"github.com/raman-vhd/video-converter/internal/lib"
)

type videoRoute struct {
    handler lib.RequestHandler
    ctrl controller.IVideoController
}

func NewVideo(
    handler lib.RequestHandler,
    ctrl controller.IVideoController,
) videoRoute{
    return videoRoute{
        handler: handler,
        ctrl: ctrl,
    }
}

func (a videoRoute) Setup() {
    api := a.handler.Echo.Group("/api")
    
    api.POST("/video", a.ctrl.CreateVideo)
    api.GET("/video/info", a.ctrl.GetVideo)
}
