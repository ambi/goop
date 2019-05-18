package usecase

import (
	"context"
	"time"

	"github.com/ambi/goop/app/config"
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
)

// Token is a use case for token issuarance.
type Token struct {
	repo db.Repository
}

// NewToken creates a Token use case.
func NewToken(repo db.Repository) *Token {
	return &Token{repo}
}

// Call processes token issuarance and returns a token or error.
func (uc *Token) Call(params *model.TokenParams) (*model.TokenResponse, *model.ClientError) {
	ctx := context.Background()

	client, _ := uc.repo.GetClient(ctx, params.ClientID)
	authzCode, _ := uc.repo.GetAuthorizationCode(ctx, params.Code)

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
