package api

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
)

type (
	Controller interface {
		GetGroup() string
		GetHandlers() []ControllerHandler
		GetMiddlewares() []middlewares.Middleware
	}
	ControllerHandler interface {
		GetMethod() string
		GetPath() string
		GetHandler() echo.HandlerFunc
	}
	Handler struct {
		Method  string
		Path    string
		Handler func(ctx echo.Context) error
	}
)

func (h *Handler) GetPath() string {
	return h.Path
}

func (h *Handler) GetHandler() echo.HandlerFunc {
	return h.Handler
}

func (h *Handler) GetMethod() string {
	return h.Method
}
