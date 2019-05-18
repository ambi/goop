package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// Consent is a controller for the OIDC consent endpoint.
type Consent struct {
	uc *usecase.Consent
}

// NewConsent creates a Consent controller.
func NewConsent(repo db.Repository) *Consent {
	return &Consent{usecase.NewConsent(repo)}
}

// Get receives a GET request to the consent endpoint, and show the consent page.
func (l *Consent) Get(c echo.Context) error {
	return c.Render(http.StatusOK, "consent", struct{}{})
}

// Post receives a POST request to the consent endpoint, and call the use case object.
func (l *Consent) Post(c echo.Context) error {
	// TODO: check CSRF token.
	// TODO
	return nil
}
