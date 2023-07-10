package lib

import "github.com/labstack/echo/v4"

type RequestHandler struct {
	Echo *echo.Echo
}

func NewRequestHandler(env Env) RequestHandler {
	e := echo.New()
	return RequestHandler{Echo: e}
}
