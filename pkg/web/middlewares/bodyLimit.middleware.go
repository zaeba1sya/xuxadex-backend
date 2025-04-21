package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type BodyLimitMiddleware struct {
	Limit string
}

func NewBodyLimitMiddleware(limit string) *BodyLimitMiddleware {
	return &BodyLimitMiddleware{
		Limit: limit,
	}
}

func (m *BodyLimitMiddleware) GetHandler() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.BodyLimit(m.Limit)
}
