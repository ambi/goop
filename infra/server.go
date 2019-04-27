package infra

import (
	"github.com/ambi/goop/app/config"
	"github.com/labstack/echo/v4"
)

// Serve は echo の Web サーバを起動する。
func Serve() {
	e := echo.New()

	dao, err := ConnectRDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Renderer = NewTemplate("resource/template/*")
	SetupSession(e)
	Route(e, dao)

	e.Logger.Fatal(e.Start(":" + config.Config.Server.ParsedURL.Port()))
}
