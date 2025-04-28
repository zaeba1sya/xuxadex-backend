package api

import (
	"context"
	"net/http/pprof"

	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
)

type profilerApi struct {
	ctx context.Context
}

func NewProfilerApi(ctx context.Context) Controller {
	return &profilerApi{
		ctx: ctx,
	}
}

func (a *profilerApi) GetGroup() string {
	return "/pprof"
}

func (a *profilerApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *profilerApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "GET",
			Path:    "/",
			Handler: a.indexHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/heap",
			Handler: a.heapHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/goroutine",
			Handler: a.goroutineHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/block",
			Handler: a.blockHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/threadcreate",
			Handler: a.threadCreateHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/cmdline",
			Handler: a.cmdlineHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/profile",
			Handler: a.profileHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/symbol",
			Handler: a.symbolHandler,
		},
		&Handler{
			Method:  "POST",
			Path:    "/symbol",
			Handler: a.symbolHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/trace",
			Handler: a.traceHandler,
		},
		&Handler{
			Method:  "GET",
			Path:    "/mutex",
			Handler: a.mutexHandler,
		},
	}
}

func (a *profilerApi) indexHandler(ctx echo.Context) error {
	pprof.Index(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) heapHandler(ctx echo.Context) error {
	pprof.Handler("/pprof/heap").ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func (a *profilerApi) goroutineHandler(ctx echo.Context) error {
	pprof.Handler("goroutine").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) blockHandler(ctx echo.Context) error {
	pprof.Handler("block").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) threadCreateHandler(ctx echo.Context) error {
	pprof.Handler("threadcreate").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) cmdlineHandler(ctx echo.Context) error {
	pprof.Cmdline(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) profileHandler(ctx echo.Context) error {
	pprof.Profile(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) symbolHandler(ctx echo.Context) error {
	pprof.Symbol(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) traceHandler(ctx echo.Context) error {
	pprof.Trace(ctx.Response().Writer, ctx.Request())
	return nil
}

func (a *profilerApi) mutexHandler(ctx echo.Context) error {
	pprof.Handler("mutex").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}
