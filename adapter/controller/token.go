package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/labstack/echo/v4"
)

// Token は Token Endpoint 用のコントローラ。
type Token struct {
	uc *usecase.Token
}

// NewToken は Token コントローラを生成する。
func NewToken(dao db.DAO) *Token {
	return &Token{usecase.NewToken(dao)}
}

// Post は Token Endpoint へのリクエストを受け付けて、実処理をユースケースに任せる。
func (t *Token) Post(c echo.Context) error {
	// TODO: Parse Authorization Header
	clientID := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	params := &model.TokenParams{
		GrantType:    c.FormValue("grant_type"),
		Code:         c.FormValue("code"),
		RedirectURI:  c.FormValue("redirect_uri"),
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	resp, err := t.uc.Call(params)

	if err != nil {
		return c.JSON(err.StatusCode, err)
	}

	return c.JSON(http.StatusOK, resp)
}
