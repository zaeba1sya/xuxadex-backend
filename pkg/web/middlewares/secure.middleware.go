package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SecureMiddleware struct{}

func NewSecureMiddleware() *SecureMiddleware {
	return &SecureMiddleware{}
}

func (m *SecureMiddleware) GetHandler() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	})
}
