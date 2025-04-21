package api

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/xuxadex/backend-mvp-main/pkg/web/middlewares"
)

type swaggerApi struct{}

func NewSwaggerApi() Controller {
	return &swaggerApi{}
}

func (*swaggerApi) GetGroup() string {
	return "/swagger"
}

func (*swaggerApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (c *swaggerApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "GET",
			Path:    "/*",
			Handler: echoSwagger.WrapHandler,
		},
	}
}
