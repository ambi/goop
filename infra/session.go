package infra

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// SetupSession sets up the session store.
func SetupSession(e *echo.Echo) {
	store := sessions.NewCookieStore([]byte("secret")) // TODO
	e.Use(session.Middleware(store))
}
