package usecase

import (
	"context"
	"time"

	"github.com/ambi/goop/app/config"
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
)

// Token はトークン生成処理のユースケース。
type Token struct {
	dao db.DAO
}

// NewToken は Token ユースケースを生成する。
func NewToken(dao db.DAO) *Token {
	return &Token{dao}
}

// Call はトークン生成処理を行って、トークンまたはエラーを返す。
func (uc *Token) Call(params *model.TokenParams) (*model.TokenResponse, *model.TokenError) {
	ctx := context.Background()

	client, _ := uc.dao.GetClient(ctx, params.ClientID)
	authzCode, _ := uc.dao.GetAuthorizationCode(ctx, params.Code)

	err := params.Valid(client, authzCode)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	resourceID := config.Config.Server.URL // TODO
	idExpiresIn := config.Config.OIDC.IDTokenExpiresIn
	acExpiresIn := config.Config.OIDC.AccessTokenExpiresIn
	idToken := model.NewIDToken(now, config.SingleOP, client.ClientID, authzCode.User, idExpiresIn, authzCode.Nonce)
	accessToken := model.NewAccessToken(now, config.SingleOP, resourceID, authzCode.User, acExpiresIn)

	resp := &model.TokenResponse{
		AccessToken: accessToken,
		IDToken:     idToken,
		TokenType:   accessToken.TokenType(),
		ExpiresIn:   acExpiresIn,
	}
	return resp, nil
}
