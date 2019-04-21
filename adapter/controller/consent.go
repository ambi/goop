package controller

import (
	"net/http"

	"github.com/ambi/goop/app/usecase"
	"github.com/ambi/goop/domain/db"
	"github.com/labstack/echo/v4"
)

// Consent は同意エンドポイント用のコントローラ。
type Consent struct {
	uc *usecase.Consent
}

// NewConsent は Consent コントローラを生成する。
func NewConsent(dao db.DAO) *Consent {
	return &Consent{usecase.NewConsent(dao)}
}

// Get は同意エンドポイントへのリクエストを受け付けて、同意画面を表示する。
func (l *Consent) Get(c echo.Context) error {
	return c.Render(http.StatusOK, "consent", struct{}{})
}

// Post は同意エンドポイントへのリクエストを受け付けて、認証処理をユースケースに任せる。
func (l *Consent) Post(c echo.Context) error {
	// TODO
	return nil
}
