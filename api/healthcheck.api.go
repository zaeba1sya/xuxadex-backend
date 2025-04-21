package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/responses"
	"github.com/xuxadex/backend-mvp-main/pkg/web/middlewares"
)

type healthcheckApi struct {
	ctx context.Context
	log logger.Logger
	db  *db.DBClient
}

func NewHealthcheckApi(ctx context.Context, log logger.Logger, db *db.DBClient) Controller {
	return &healthcheckApi{
		ctx: ctx,
		log: log,
		db:  db,
	}
}

func (a *healthcheckApi) GetGroup() string {
	return "/"
}

func (a *healthcheckApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *healthcheckApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "GET",
			Path:    "healthcheck",
			Handler: a.healthcheck,
		},
	}
}

// Healthcheck godoc
// @Summary      Healthcheck
// @Description  Healthcheck handler
// @Tags         system
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=string}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /healthcheck [get]
func (a *healthcheckApi) healthcheck(ctx echo.Context) error {
	row := a.db.GetClient().QueryRow("select now() as t")
	if row.Err() != nil {
		a.log.Info("Failed to get current time")
	}
	var t time.Time
	err := row.Scan(&t)
	if err != nil {
		a.log.Error("Failed to scan current time", err)
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, fmt.Sprintf("Database Time: %s", t.Format("02.01.2006 15:04:05")), true)
}
