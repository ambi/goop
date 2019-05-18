package infra

import (
	"github.com/ambi/goop/adapter/gateway"
	"github.com/ambi/goop/app/config"
	"github.com/labstack/echo/v4"
)

// Serve starts a web application.
func Serve() {
	e := echo.New()

	rdb, err := OpenRDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	repo, err := gateway.NewSQL(rdb)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Renderer = NewTemplate("resource/template/*")
	SetupSession(e)
	Route(e, repo)

	e.Logger.Fatal(e.Start(":" + config.Config.Server.ParsedURL.Port()))
}
