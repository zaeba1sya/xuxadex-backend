package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/responses"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
)

type systemApi struct {
	ctx context.Context
	log logger.Logger
	db  *db.DBClient
}

func NewSystemApi(ctx context.Context, log logger.Logger, db *db.DBClient) Controller {
	return &systemApi{
		ctx: ctx,
		log: log,
		db:  db,
	}
}

func (a *systemApi) GetGroup() string {
	return "/"
}

func (a *systemApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *systemApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "GET",
			Path:    "healthcheck",
			Handler: a.healthcheck,
		},
		&Handler{
			Method:  "GET",
			Path:    "static/*",
			Handler: echo.WrapHandler(http.StripPrefix("/api/v1/static", http.FileServer(http.Dir("static")))),
		},
	}
}

// System godoc
// @Summary      System
// @Description  System handler
// @Tags         system
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=string}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /healthcheck [get]
func (a *systemApi) healthcheck(ctx echo.Context) error {
	anchor := "healthcheck"
	row := a.db.GetClient().QueryRow("select now() as t")
	if row.Err() != nil {
		a.log.Errorf("[%s] error: failed to get current time", anchor)
	}
	var t time.Time
	err := row.Scan(&t)
	if err != nil {
		a.log.Errorf("[%s] error: failed to scan current time", anchor)
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, fmt.Sprintf("database time: %s", t.Format("02.01.2006 15:04:05")), true)
}
