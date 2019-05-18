package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/labstack/echo/v4"
)

// Token is a controller for the OAuth 2.0 token endpoint.
type Token struct {
	uc *usecase.Token
}

// NewToken creates a Token controller.
func NewToken(repo db.Repository) *Token {
	return &Token{usecase.NewToken(repo)}
}

// Post receives a request to the token endpoint, and call the use case object.
func (t *Token) Post(c echo.Context) error {
	// No CSRF token.
	// TODO: Parse Authorization Header
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	params := &model.TokenParams{
		GrantType:    c.FormValue("grant_type"),
		Code:         c.FormValue("code"),
		RedirectURI:  c.FormValue("redirect_uri"),
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	resp, err := t.uc.Call(params)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}

	return c.JSON(http.StatusOK, resp)
}
