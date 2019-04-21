package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// Login はログインエンドポイント用のコントローラ。
type Login struct {
	uc *usecase.Login
}

// NewLogin は Login コントローラを生成する。
func NewLogin(dao db.DAO) *Login {
	return &Login{usecase.NewLogin(dao)}
}

// Get はログインエンドポイントへのリクエストを受け付けて、ログイン画面を表示する。
func (l *Login) Get(c echo.Context) error {
	return c.Render(http.StatusOK, "login", struct{}{})
}

// Post はログインエンドポイントへのリクエストを受け付けて、認証処理をユースケースに任せる。
func (l *Login) Post(c echo.Context) error {
	loginID := c.FormValue("login_id")
	password := c.FormValue("password")
	sess := GetSession(c)

	_, err := l.uc.Call(sess, loginID, password)
	SaveSession(c, sess)

	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}

	return c.Redirect(http.StatusFound, "/consent") // TODO
}
