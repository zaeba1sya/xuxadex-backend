package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/internal/user"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/responses"
	"github.com/xuxadex/backend-mvp-main/pkg/types"

	"github.com/xuxadex/backend-mvp-main/pkg/web/middlewares"
)

type authApi struct {
	ctx        context.Context
	log        logger.Logger
	repository *user.UserService
}

func NewAuthApi(ctx context.Context, log logger.Logger, db *db.DBClient) Controller {
	return &authApi{
		ctx:        ctx,
		log:        log,
		repository: user.NewUserService(db, log),
	}
}

func (a *authApi) GetGroup() string {
	return "/"
}

func (a *authApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (a *authApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{
		&Handler{
			Method:  "GET",
			Path:    "auth",
			Handler: a.auth,
		},
	}
}

// Auth godoc
// @Summary      Auth
// @Description  Auth handler
// @Tags         user
// @Param        wallet    query     string  false  "0x..."  string
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=user.UserEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /auth [get]
func (a *authApi) auth(ctx echo.Context) error {
	anchor := "auth"
	a.log.Infof("[%s] Request received", anchor)

	wallet := types.Wallet(ctx.QueryParam("wallet"))

	if err := wallet.Validate(); err != nil {
		a.log.Infof("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, err.Error(), false)
	}

	user, err := a.repository.Authenticate(a.ctx, wallet)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, user, true)
}
