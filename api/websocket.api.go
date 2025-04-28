package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/responses"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/sockets"
)

type websocketApi struct {
	ctx        context.Context
	socketPool *sockets.SocketsPool
	log        logger.Logger
	upgrader   *websocket.Upgrader
}

func NewWebsocketApi(ctx context.Context, socketPool *sockets.SocketsPool, log logger.Logger) Controller {
	return &websocketApi{
		ctx:        ctx,
		socketPool: socketPool,
		log:        log,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (a *websocketApi) GetGroup() string {
	return "/ws"
}

func (a *websocketApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *websocketApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "GET",
			Path:    "/activity/:activityID",
			Handler: a.activity,
		},
		&Handler{
			Method:  "GET",
			Path:    "/activity/send",
			Handler: a.foo,
		},
	}
}

func (a *websocketApi) foo(ctx echo.Context) error {
	err := a.socketPool.SendToAll("4847209f-3ade-47ff-a28b-cdeaeaa463b2", "Hello World!")
	if err != nil {
		a.log.Info(err.Error())
	}
	return nil
}

// Websocket godoc
// @Summary      Websocket
// @Description  Websocket handler
// @Tags         websocket
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=string}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /ws/activity/{activityID} [get]
func (a *websocketApi) activity(ctx echo.Context) error {
	anchor := "websocket activity"

	activityID := ctx.Param("activityID")
	if activityID == "" {
		a.log.Errorf("[%s] connection error: %s", anchor, "activityID is required")
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, "activityID is required", false)
	}

	conn, err := a.upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		a.log.Errorf("[%s] connection error: %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}
	conn.SetReadDeadline(time.Time{})
	conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})
	defer conn.Close()

	connID := a.socketPool.AddConnection(activityID, conn)
	defer a.socketPool.RemoveConnection(activityID, connID)

	a.log.Infof("[%s] connection established: %s", anchor, connID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			a.log.Infof("[%s] %s (%s)", anchor, err.Error(), connID)
			break
		}
	}

	return nil
}
