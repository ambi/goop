package usecase

import (
	"context"
	"net/http"

	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/labstack/gommon/log"
)

// Authorize is a use case for authorization.
type Authorize struct {
	repo db.Repository
}

// NewAuthorize creates an Authorize use case.
func NewAuthorize(repo db.Repository) *Authorize {
	return &Authorize{repo}
}

// Call processes authorization and returns an authorization code or error.
func (uc *Authorize) Call(params *model.AuthorizeParams, loginID string) (authzResponse *model.AuthorizeResponse, authzErr *model.AuthorizeError, loginRequired bool) {
	ctx := context.Background()

	client, getErr := uc.repo.GetClient(ctx, params.ClientID)
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
	user, getErr := uc.repo.GetUser(ctx, loginID)
	if getErr != nil {
		log.Error(getErr)
		loginRequired = true
		return
	}
	authzCode, createErr := uc.repo.CreateAuthorizationCode(ctx, user)
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
