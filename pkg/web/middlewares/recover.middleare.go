package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RecoverMiddleware struct{}

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

func (m *RecoverMiddleware) GetHandler() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.Recover()
}
