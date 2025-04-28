package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/internal/game"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/repository"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/responses"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
)

type gameApi struct {
	ctx     context.Context
	log     logger.Logger
	service *game.GameService
}

func NewGameApi(ctx context.Context, log logger.Logger, db *db.DBClient) Controller {
	return &gameApi{
		ctx:     ctx,
		log:     log,
		service: game.NewGameService(db),
	}
}

func (a *gameApi) GetGroup() string {
	return "/game"
}

func (a *gameApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *gameApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "POST",
			Path:    "",
			Handler: a.create,
		},
		&Handler{
			Method:  "GET",
			Path:    "s",
			Handler: a.getAll,
		},
		&Handler{
			Method:  "GET",
			Path:    "/:id",
			Handler: a.getByID,
		},
		&Handler{
			Method:  "DELETE",
			Path:    "/:id",
			Handler: a.delete,
		},
	}
}

// Get All Games godoc
// @Summary      Get All Games
// @Description  Get All Games handler
// @Tags         game
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=[]game.GameEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /games [get]
func (a *gameApi) getAll(ctx echo.Context) error {
	anchor := "get all games"
	a.log.Infof("[%s] request received", anchor)

	queryOpts := repository.ParseQueryOpts(ctx)

	games, err := a.service.GetAll(a.ctx, queryOpts)
	if err != nil {
		a.log.Errorf("[%s] error: %v", anchor, err)
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, games, true)
}

// Game create godoc
// @Summary      Game create
// @Description  Game create handler
// @Tags         game
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=game.GameEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /game [post]
func (a *gameApi) create(ctx echo.Context) error {
	anchor := "create game"
	a.log.Infof("[%s] request received", anchor)

	data := &game.GameCreateDTO{}

	if err := ctx.Bind(data); err != nil {
		a.log.Errorf("[%s] error: %v", anchor, err)
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, err.Error(), false)
	}
	if err := data.Validate(); err != nil {
		a.log.Errorf("[%s] error: %v", anchor, err)
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, err.Error(), false)
	}

	games, err := a.service.Create(a.ctx, data)
	if err != nil {
		a.log.Errorf("[%s] error: %v", anchor, err)
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, games, true)
}

// Get Game By ID godoc
// @Summary      Get Game By ID
// @Description  Get Game By ID handler
// @Tags         game
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=game.GameEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /game/{id} [get]
func (a *gameApi) getByID(ctx echo.Context) error {
	anchor := "get by id game"
	a.log.Infof("[%s] request received", anchor)

	id := ctx.Param("id")

	if id == "" {
		a.log.Errorf("[%s] error: %s", anchor, "id is required")
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, "id is required", false)
	}

	game, err := a.service.GetByID(a.ctx, id)
	if err != nil {
		a.log.Errorf("[%s] error: %v", anchor, err)
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, game, true)
}

// Delete Game By ID godoc
// @Summary      Delete Game By ID
// @Description  Delete Game By ID handler
// @Tags         game
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=string}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /game/{id} [delete]
func (a *gameApi) delete(ctx echo.Context) error {
	anchor := "delete game"
	a.log.Infof("[%s] request received", anchor)

	id := ctx.Param("id")

	if id == "" {
		a.log.Errorf("[%s] error: %s", anchor, "id is required")
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, "id is required", false)
	}

	err := a.service.Delete(a.ctx, id)
	if err != nil {
		a.log.Errorf("[%s] error: %v", anchor, err)
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, "game deleted successfully", true)
}
