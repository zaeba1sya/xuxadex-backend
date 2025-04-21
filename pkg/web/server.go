package web

import (
	"context"
	"fmt"

	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/web/middlewares"

	"github.com/xuxadex/backend-mvp-main/api"
	"github.com/xuxadex/backend-mvp-main/config"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	cfg    *config.Config
	log    logger.Logger
	client *echo.Echo
}

func NewWebServer(cfg *config.Config, log logger.Logger) *WebServer {
	return &WebServer{
		cfg:    cfg,
		log:    log,
		client: echo.New(),
	}
}

func (w *WebServer) RegisterRoutes(routes []api.Controller) {
	for _, route := range routes {
		mainGroup := w.client.Group("/api/v1")
		group := mainGroup.Group(route.GetGroup())

		for _, handler := range route.GetHandlers() {
			switch handler.GetMethod() {
			case "GET":
				group.GET(handler.GetPath(), handler.GetHandler())
			case "POST":
				group.POST(handler.GetPath(), handler.GetHandler())
			case "PUT":
				group.PUT(handler.GetPath(), handler.GetHandler())
			case "PATCH":
				group.PATCH(handler.GetPath(), handler.GetHandler())
			case "DELETE":
				group.DELETE(handler.GetPath(), handler.GetHandler())
			default:
				w.log.Error("Unsupported method")
			}
		}
	}
}

func (w *WebServer) RegisterMiddlewares(middlewares []middlewares.Middleware) {
	for _, middlware := range middlewares {
		w.client.Use(middlware.GetHandler())
	}
}

func (w *WebServer) Listen() {
	w.client.HideBanner = true
	w.client.HidePort = true
	err := w.client.Start(fmt.Sprintf("%s:%d", w.cfg.Server.Host, w.cfg.Server.Port))

	if err != nil {
		w.log.Fatal(err.Error())
	}
}

func (s *WebServer) Release(ctx context.Context) error {
	return s.client.Shutdown(ctx)
}
