package infra

import (
	"github.com/ambi/goop/adapter/controller"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// Route はルーティングを行う。
func Route(e *echo.Echo, dao db.DAO) {
	openIDConfiguration := controller.NewOpenIDConfiguration(dao)
	e.GET("/.well-known/openid-configuration", openIDConfiguration.Get)

	jwks := controller.NewJWKS(dao)
	e.GET("/jwks", jwks.Get)

	authorize := controller.NewAuthorize(dao)
	e.GET("/authorize", authorize.Get)
	e.POST("/authorize", authorize.Post)

	token := controller.NewToken(dao)
	e.POST("/token", token.Post)

	login := controller.NewLogin(dao)
	e.GET("/login", login.Get)
	e.POST("/login", login.Post)

	consent := controller.NewConsent(dao)
	e.GET("/consent", consent.Get)
	e.POST("/consent", consent.Post)
}
