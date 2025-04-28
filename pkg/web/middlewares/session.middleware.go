package middlewares

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/config"
)

type SessionMiddleware struct {
	cfg *config.Config
}

func NewSessionMiddleware(cfg *config.Config) *SessionMiddleware {
	return &SessionMiddleware{
		cfg: cfg,
	}
}

func (m *SessionMiddleware) GetHandler() func(next echo.HandlerFunc) echo.HandlerFunc {
	return session.Middleware(sessions.NewCookieStore([]byte(m.cfg.Server.Session.Secret)))
}
