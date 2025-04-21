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
			Path:    "system",
			Handler: a.system,
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
// @Router       /system [get]
func (a *systemApi) system(ctx echo.Context) error {
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
