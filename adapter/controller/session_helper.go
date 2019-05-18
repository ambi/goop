package controller

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// GetSession gets session data.
func GetSession(c echo.Context) *sessions.Session {
	sess, err := session.Get("session", c) // TODO
	if err != nil {
		c.Logger().Errorf("can't get session: error=%v\n", err)
		return nil
	}
	return sess
}

// SaveSession saves session data.
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
