package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type StaticMiddleware struct{}

func NewStaticMiddleware() *StaticMiddleware {
	return &StaticMiddleware{}
}

func (m *StaticMiddleware) GetHandler() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
	})
}
