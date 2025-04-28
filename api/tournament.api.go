package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/internal/tournament"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/repository"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/responses"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
)

type tournamentApi struct {
	ctx     context.Context
	log     logger.Logger
	service *tournament.TournamentService
}

func NewTournamentApi(ctx context.Context, log logger.Logger, db *db.DBClient) Controller {
	return &tournamentApi{
		ctx:     ctx,
		log:     log,
		service: tournament.NewTournamentService(db, log),
	}
}

func (a *tournamentApi) GetGroup() string {
	return "/tournament"
}

func (a *tournamentApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *tournamentApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "GET",
			Path:    "s",
			Handler: a.getTournaments,
		},
		&Handler{
			Method:  "GET",
			Path:    "/dashboard",
			Handler: a.dashboard,
		},
		&Handler{
			Method:  "GET",
			Path:    "/:id",
			Handler: a.getByID,
		},
		&Handler{
			Method:  "POST",
			Path:    "",
			Handler: a.createTournament,
		},
		&Handler{
			Method:  "POST",
			Path:    "/join/:id",
			Handler: a.joinTournament,
		},
		&Handler{
			Method:  "GET",
			Path:    "/statuses",
			Handler: a.getTournamentStatuses,
		},
		&Handler{
			Method:  "GET",
			Path:    "/randomize",
			Handler: a.randomizeDates,
		},
	}
}

// Tournaments Dashboard godoc
// @Summary      Tournaments Dashboard
// @Description  Tournaments Dashboard handler
// @Tags         tournament
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=tournament.TournamentDashboardDTO}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /tournament/dashboard [get]
func (a *tournamentApi) dashboard(ctx echo.Context) error {
	anchor := "tournament dashboard"
	a.log.Infof("[%s] request received", anchor)

	dashboard, err := a.service.GetDashboard(a.ctx)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, dashboard, true)
}

// All Tournaments godoc
// @Summary      All Tournaments
// @Description  All Tournaments handler
// @Tags         tournament
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "1"          string
// @Param        limit     query     string  false  "10"         string
// @Param        sort      query     string  false  "id:ASC"     string
// @Param        filter    query     string  false  "count<100"  string
// @Success      200  {object}  responses.Response{data=tournament.TournamentBaseEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /tournaments [get]
func (a *tournamentApi) getTournaments(ctx echo.Context) error {
	anchor := "get tournaments"
	a.log.Infof("[%s] request received", anchor)

	queryParams := repository.ParseQueryOpts(ctx)

	tournaments, err := a.service.GetAll(a.ctx, queryParams)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, tournaments, true)
}

// Create Tournaments godoc
// @Summary      Create Tournaments
// @Description  Create Tournaments handler
// @Tags         tournament
// @Accept       json
// @Produce      json
// @Param        tournament body tournament.TournamentCreateDTO true "Tournament create object"
// @Success      200  {object}  responses.Response{data=tournament.TournamentBaseEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /tournament [post]
func (a *tournamentApi) createTournament(ctx echo.Context) error {
	anchor := "create tournament"
	a.log.Infof("[%s] request received", anchor)

	data := &tournament.TournamentCreateDTO{}
	if err := ctx.Bind(data); err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, err.Error(), false)
	}
	tournament, err := a.service.CreateWithRelations(a.ctx, data)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusCreated, tournament, true)
}

func (a *tournamentApi) joinTournament(ctx echo.Context) error {
	anchor := "join tournament"
	a.log.Infof("[%s] request received", anchor)

	data := &tournament.TournamentJoinDTO{}

	data.TournamentID = ctx.Param("id")
	if data.TournamentID == "" {
		a.log.Errorf("[%s] %s", anchor, "tournament ID is required")
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, "tournament ID is required", false)
	}

	if err := ctx.Bind(data); err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, "joined tournament successfully", true)
}

// Tournament Get By ID godoc
// @Summary      Tournament Get By ID
// @Description  Tournament Get By ID handler
// @Tags         tournament
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Tournament ID"
// @Success      200  {object}  responses.Response{data=tournament.TournamentFullEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /tournament/{id} [get]
func (a *tournamentApi) getByID(ctx echo.Context) error {
	anchor := "get tournament by id"
	a.log.Infof("[%s] request received", anchor)

	id := ctx.Param("id")

	if id == "" {
		a.log.Errorf("[%s] %s", anchor, "id is required")
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, "id is required", false)
	}

	tournament, err := a.service.GetById(a.ctx, id)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, tournament, true)
}

// Tournament Statuses godoc
// @Summary      Tournament Statuses
// @Description  Tournament Statuses handler
// @Tags         tournament
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=[]tournament.TournamentStatusEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /tournament/statuses [get]
func (a *tournamentApi) getTournamentStatuses(ctx echo.Context) error {
	anchor := "get tournament types"
	a.log.Infof("[%s] request received", anchor)

	statuses, err := a.service.GetStatuses(a.ctx)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, int(http.StatusInternalServerError), err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, int(http.StatusOK), statuses, true)
}

// Randomize Dates godoc
// @Summary      Randomize Dates
// @Description  Randomize Dates handler
// @Tags         tournament
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=boolean}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /tournament/randomize [get]
func (a *tournamentApi) randomizeDates(ctx echo.Context) error {
	anchor := "randomize tournaments dates"
	a.log.Infof("[%s] request received", anchor)

	err := a.service.RandomizeDates(a.ctx)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, int(http.StatusInternalServerError), err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, int(http.StatusOK), "dates randomized successfully", true)
}
