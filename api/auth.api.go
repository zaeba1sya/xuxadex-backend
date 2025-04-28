package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/internal/user"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/responses"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/types"

	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
	sessionmgr "gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/session"
)

type authApi struct {
	ctx     context.Context
	log     logger.Logger
	service *user.UserService
}

func NewAuthApi(ctx context.Context, log logger.Logger, db *db.DBClient) Controller {
	return &authApi{
		ctx:     ctx,
		log:     log,
		service: user.NewUserService(db, log),
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
		&Handler{
			Method:  "GET",
			Path:    "me",
			Handler: a.me,
		},
		&Handler{
			Method:  "GET",
			Path:    "logout",
			Handler: a.logout,
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

	user, err := a.service.Authenticate(a.ctx, wallet)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	if err := sessionmgr.CreateSession(ctx, user.ID); err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, user, true)
}

// Get User Data From Session godoc
// @Summary      Get User Data From Session
// @Description  Get User Data From Session handler
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=user.UserEntity}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /me [get]
func (a *authApi) me(ctx echo.Context) error {
	anchor := "get session user data"
	a.log.Infof("[%s] Request received", anchor)

	userID, err := sessionmgr.GetIdFromSession(ctx)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusBadRequest, err.Error(), false)
	}

	user, err := a.service.GetById(a.ctx, userID)
	if err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, user, true)
}

// Logout godoc
// @Summary      Logout
// @Description  Logout handler
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  responses.Response{data=bool}
// @Failure      400,500  {object}  responses.Response{data=string}
// @Router       /logout [get]
func (a *authApi) logout(ctx echo.Context) error {
	anchor := "logout session"
	a.log.Infof("[%s] Request received", anchor)

	if err := sessionmgr.InvalidateSession(ctx); err != nil {
		a.log.Errorf("[%s] %s", anchor, err.Error())
		return responses.NewApplicationResponse(ctx, http.StatusInternalServerError, err.Error(), false)
	}

	return responses.NewApplicationResponse(ctx, http.StatusOK, true, true)
}
