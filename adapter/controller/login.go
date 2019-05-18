package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// Login is a controller for the login endpoint.
type Login struct {
	uc *usecase.Login
}

// NewLogin creates a Login controller.
func NewLogin(repo db.Repository) *Login {
	return &Login{usecase.NewLogin(repo)}
}

// Get receives a GET request to the login endpoint, and show the login page.
func (l *Login) Get(c echo.Context) error {
	return c.Render(http.StatusOK, "login", struct{}{})
}

// Post receives a POST request to the login endpoint, and call the use case object.
func (l *Login) Post(c echo.Context) error {
	// TODO: check CSRF token.
	loginID := c.FormValue("login_id")
	password := c.FormValue("password")
	sess := GetSession(c)

	_, err := l.uc.Call(sess, loginID, password)
	SaveSession(c, sess)

	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}

	return c.Redirect(http.StatusFound, "/consent") // TODO
}
