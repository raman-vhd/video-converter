package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raman-vhd/video-converter/internal/service"
)

type ITemplateController interface{
    Action(ctx echo.Context) error
}

type templateController struct {
    svc service.ITemplateService
}

func NewTemplate(
    svc service.ITemplateService,
) ITemplateController{
    return templateController{
        svc: svc,
    }
}

func (c templateController) Action(ctx echo.Context) error {
    return ctx.JSON(http.StatusNotImplemented, "not implemented")
}
