package controller

import (
	"net/http"

	"github.com/ambi/goop/app/config"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// OpenIDConfiguration is a controller for the OP configuration endpoint.
type OpenIDConfiguration struct{}

// NewOpenIDConfiguration creates an OpenIDConfiguration controller.
func NewOpenIDConfiguration(_ db.Repository) *OpenIDConfiguration {
	return &OpenIDConfiguration{}
}

// Get receives a GET request to the OP configuration endpoint, and returns the OP configuration metadata.
func (oc *OpenIDConfiguration) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, config.SingleOP)
}
