package infra

import (
	"github.com/ambi/goop/adapter/controller"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// Route sets up routing.
func Route(e *echo.Echo, repo db.Repository) {
	openIDConfiguration := controller.NewOpenIDConfiguration(repo)
	e.GET("/.well-known/openid-configuration", openIDConfiguration.Get)

	jwks := controller.NewJWKS(repo)
	e.GET("/jwks", jwks.Get)

	authorize := controller.NewAuthorize(repo)
	e.GET("/authorize", authorize.Get)
	e.POST("/authorize", authorize.Post)

	token := controller.NewToken(repo)
	e.POST("/token", token.Post)

	login := controller.NewLogin(repo)
	e.GET("/login", login.Get)
	e.POST("/login", login.Post)

	consent := controller.NewConsent(repo)
	e.GET("/consent", consent.Get)
	e.POST("/consent", consent.Post)
}
