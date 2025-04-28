package sessionmgr

import (
	"errors"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetIdFromSession(ctx echo.Context) (string, error) {
	sess, err := session.Get("session", ctx)
	if err != nil {
		return "", err
	}

	id, ok := sess.Values["id"].(string)
	if !ok {
		return "", errors.New("not authorized")
	}

	return id, nil
}

func CreateSession(ctx echo.Context, id string) error {
	sess, err := session.Get("session", ctx)
	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values["id"] = id

	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		return err
	}

	return nil
}

func InvalidateSession(ctx echo.Context) error {
	sess, err := session.Get("session", ctx)
	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		return err
	}

	return nil
}
