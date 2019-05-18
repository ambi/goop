package db

import (
	"context"
	"errors"
	"time"

	"github.com/ambi/goop/domain/model"
)

// Repository is a repository for all data in DB.
type Repository interface {
	Close() error

	CreateAuthorizationCode(ctx context.Context, user *model.User) (*model.AuthorizationCode, error)
	GetAuthorizationCode(ctx context.Context, code string) (*model.AuthorizationCode, error)

	CreateClient(ctx context.Context, name string) (*model.Client, error)
	GetClient(ctx context.Context, clientID string) (*model.Client, error)

	CreateUser(ctx context.Context, loginID string) (*model.User, error)
	GetUser(ctx context.Context, loginID string) (*model.User, error)

	CreateRevocation(ctx context.Context, token string, expiresAt time.Time) error
	GetRevocation(ctx context.Context, token string) (bool, error)
}

var (
	// ErrNotFound is an error when data cannot be found.
	ErrNotFound = errors.New("db.ErrNotFound")

	// ErrNotSaved is an error when data cannot be saved.
	ErrNotSaved = errors.New("db.ErrNotSaved")
)
