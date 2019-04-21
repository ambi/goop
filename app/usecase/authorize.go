package usecase

import (
	"context"
	"net/http"

	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/labstack/gommon/log"
)

// Authorize は認可処理のユースケース。
type Authorize struct {
	dao db.DAO
}

// NewAuthorize は Authorize ユースケースを生成する。
func NewAuthorize(dao db.DAO) *Authorize {
	return &Authorize{dao}
}

// Call は認可処理を行って、認可コードまたはエラーを返す。
func (uc *Authorize) Call(params *model.AuthorizeParams, loginID string) (authzResponse *model.AuthorizeResponse, authzErr *model.AuthorizeError, loginRequired bool) {
	ctx := context.Background()

	client, getErr := uc.dao.GetClient(ctx, params.ClientID)
	if getErr != nil {
		log.Error(getErr)
	}
	err := params.Valid(client)
	if err != nil {
		authzErr = err
		return
	}

	if loginID == "" {
		loginRequired = true
		return
	}
	user, getErr := uc.dao.GetUser(ctx, loginID)
	if getErr != nil {
		log.Error(getErr)
		loginRequired = true
		return
	}
	authzCode, createErr := uc.dao.CreateAuthorizationCode(ctx, user)
	if createErr != nil {
		log.Error(createErr)
		authzErr = &model.AuthorizeError{
			StatusCode: http.StatusFound,
			Message:    "server_error",
			State:      params.State,
		}
		return
	}

	authzResponse = &model.AuthorizeResponse{
		Code:  authzCode.Code,
		State: params.State,
	}
	return
}
