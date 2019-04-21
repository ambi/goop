package controller

import (
	"net/http"

	"github.com/ambi/goop/app/config"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// OpenIDConfiguration は OP Configuration Endpoint 用のコントローラ。
type OpenIDConfiguration struct{}

// NewOpenIDConfiguration は OpenIDConfiguration コントローラを生成する。
func NewOpenIDConfiguration(_ db.DAO) *OpenIDConfiguration {
	return &OpenIDConfiguration{}
}

// Get は OP Configuration Endpoint へのリクエストを受け取って、OP Configuration Metdata JSON を返す。
func (oc *OpenIDConfiguration) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, config.SingleOP)
}
