package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/internal/match"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/repository"
	"github.com/xuxadex/backend-mvp-main/pkg/responses"

	server_errors "github.com/xuxadex/backend-mvp-main/pkg/web/errors"
	"github.com/xuxadex/backend-mvp-main/pkg/web/middlewares"
)

type matchApi struct {
	ctx     context.Context
	log     logger.Logger
	service *match.MatchService
}

func NewMatchApi(ctx context.Context, log logger.Logger, db *db.DBClient) Controller {
	return &matchApi{
		ctx:     ctx,
		log:     log,
		service: match.NewMatchService(db),
	}
}

func (a *matchApi) GetGroup() string {
	return "/match"
}

func (a *matchApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *matchApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "POST",
			Path:    "/quick",
			Handler: a.quickMatchCreate,
		},
		&Handler{
			Method:  "GET",
			Path:    "/all",
			Handler: a.getMatches,
		},
	}
}

// Get Quick Matches godoc
// @Summary      Get Quick Matches
// @Description  Get Quick Matches handler
// @Tags         match
// @Param        page      query     string  false  "1"          string
// @Param        limit     query     string  false  "10"         string
// @Param        sort      query     string  false  "id:ASC"     string
// @Param        filter    query     string  false  "count<100"  string
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=[]match.MatchEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /match/all [get]
func (a *matchApi) getMatches(ctx echo.Context) error {
	anchor := "get quick matches"
	a.log.Infof("[%s] Request received", anchor)

	queryParams := repository.ParseQueryOpts(ctx)

	matches, err := a.service.GetAll(a.ctx, queryParams)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, matches, true)
}

// Create Quick Match godoc
// @Summary      Create Quick Match
// @Description  Create Quick Match handler
// @Tags         match
// @Param input body match.QuickMatchCreateDTO true "quick match create data"
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=match.MatchEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /match/quick [post]
func (a *matchApi) quickMatchCreate(ctx echo.Context) error {
	anchor := "quick match create"
	a.log.Infof("[%s] Request received", anchor)

	data := &match.QuickMatchCreateDTO{}

	if err := ctx.Bind(data); err != nil {
		a.log.Errorf("[%s] %s", anchor, server_errors.BindError.Error())
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, server_errors.BindError.Error(), false)
	}

	match, err := a.service.CreateQuickMatch(a.ctx, data)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, match, true)
}
