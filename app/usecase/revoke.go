package usecase

import (
	"context"

	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
)

// Revoke is a use case for token revocation.
type Revoke struct {
	repo db.Repository
}

// NewRevoke creates a Revoke use case.
func NewRevoke(repo db.Repository) *Revoke {
	return &Revoke{repo}
}

// Call processes token revocation and returns nil or an error.
func (uc *Revoke) Call(params *model.RevokeParams) *model.ClientError {
	ctx := context.Background()

	client, _ := uc.repo.GetClient(ctx, params.ClientID)

	err := params.Valid(client)
	if err != nil {
		return err
	}

	// TODO: validate the token and revoke it or return an error if it's invalid.

	return nil
}
