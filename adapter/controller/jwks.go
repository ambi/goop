package controller

import (
	"net/http"

	"github.com/ambi/goop/app/config"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// JWKS は JWK Set Endpoint 用のコントローラ。
type JWKS struct{}

// NewJWKS は JWKS コントローラを生成する。
func NewJWKS(_ db.DAO) *JWKS {
	return &JWKS{}
}

// Get は JWKS へのリクエストを受け取って、JWK Set を返す。
func (j *JWKS) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, config.SingleOP.PublicJWKS)
}
