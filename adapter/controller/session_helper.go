package controller

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// GetSession はセッション情報を取得する。
func GetSession(c echo.Context) *sessions.Session {
	sess, err := session.Get("session", c) // TODO
	if err != nil {
		c.Logger().Errorf("can't get session: error=%v\n", err)
		return nil
	}
	return sess
}

// SaveSession はセッション情報を保存する。
func SaveSession(c echo.Context, sess *sessions.Session) {
	if sess == nil {
		return
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // TODO
		HttpOnly: true,
	}
	sess.Save(c.Request(), c.Response())
}
