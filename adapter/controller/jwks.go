package controller

import (
	"net/http"

	"github.com/ambi/goop/app/config"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// JWKS is a controller for the OIDC JWK Set endpoint.
type JWKS struct{}

// NewJWKS creates a controller for JWKS.
func NewJWKS(_ db.Repository) *JWKS {
	return &JWKS{}
}

// Get receives a GET request to the JWKS endpoint, and returns the JWK Set.
func (j *JWKS) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, config.SingleOP.PublicJWKS)
}
