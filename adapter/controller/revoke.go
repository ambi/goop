package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/labstack/echo/v4"
)

// Revoke is a controller for the OAuth 2.0 revocation endpoint.
type Revoke struct {
	uc *usecase.Revoke
}

// NewRevoke creates a Revoke controller.
func NewRevoke(repo db.Repository) *Revoke {
	return &Revoke{usecase.NewRevoke(repo)}
}

// Post receives a request to the revocation endpoint, and call the use case object.
func (t *Revoke) Post(c echo.Context) error {
	// No CSRF token.
	// TODO: Parse Authorization Header
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	params := &model.RevokeParams{
		Token:         c.FormValue("token"),
		TokenTypeHint: c.FormValue("token_type_hint"),
		ClientID:      clientID,
		ClientSecret:  clientSecret,
	}

	err := t.uc.Call(params)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}

	return c.JSON(http.StatusOK, struct{ Message string }{"OK"})
}
