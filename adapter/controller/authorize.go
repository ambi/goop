package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/labstack/echo/v4"
)

// Authorize is a controller for the OAuth 2.0 authorization endpoint.
type Authorize struct {
	uc *usecase.Authorize
}

// NewAuthorize creates a new Authorize controller.
func NewAuthorize(repo db.Repository) *Authorize {
	return &Authorize{usecase.NewAuthorize(repo)}
}

// Get receive a GET request to the authorization endpoint, and call the use case object.
func (a *Authorize) Get(c echo.Context) error {
	sess := GetSession(c)
	loginID, _ := sess.Values["login_id"].(string)

	params := &model.AuthorizeParams{
		Scope:        c.QueryParam("scope"),
		ResponseType: c.QueryParam("response_type"),
		ClientID:     c.QueryParam("client_id"),
		RedirectURI:  c.QueryParam("redirect_uri"),
		State:        c.QueryParam("state"),
		ResponseMode: c.QueryParam("response_mode"),
		Nonce:        c.QueryParam("nonce"),
	}

	return a.callUsecase(c, params, loginID)
}

// Post receive a POST request to the authorization endpoint, and call the use case object.
func (a *Authorize) Post(c echo.Context) error {
	// No CSRF token.
	sess := GetSession(c)
	loginID, _ := sess.Values["login_id"].(string)

	params := &model.AuthorizeParams{
		Scope:        c.FormValue("scope"),
		ResponseType: c.FormValue("response_type"),
		ClientID:     c.FormValue("client_id"),
		RedirectURI:  c.FormValue("redirect_uri"),
		State:        c.FormValue("state"),
		ResponseMode: c.FormValue("response_mode"),
		Nonce:        c.FormValue("nonce"),
	}

	return a.callUsecase(c, params, loginID)
}

func (a *Authorize) callUsecase(c echo.Context, params *model.AuthorizeParams, loginID string) error {
	resp, err, loginRequired := a.uc.Call(params, loginID)
	if loginRequired {
		return c.Redirect(http.StatusFound, "/op/login") // TODO: don't use magic number.
	}
	if err != nil {
		if err.StatusCode == http.StatusFound {
			location := params.RedirectURI + "?" + err.ToQuery() // TODO: RedirectURI may contain query parameters.
			return c.Redirect(http.StatusFound, location)
		}

		data := struct {
			StatusCode int
			Error      string
		}{
			err.StatusCode,
			err.Message,
		}
		return c.Render(err.StatusCode, "authorize_error", data)
	}

	location := params.RedirectURI + "?" + resp.ToQuery()
	return c.Redirect(http.StatusFound, location)
}
