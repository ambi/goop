package db

import (
	"context"
	"errors"

	"github.com/ambi/goop/domain/model"
)

// DAO は DAO インタフェース。
type DAO interface {
	Close() error

	CreateAuthorizationCode(ctx context.Context, user *model.User) (*model.AuthorizationCode, error)
	GetAuthorizationCode(ctx context.Context, code string) (*model.AuthorizationCode, error)

	CreateClient(ctx context.Context, name string) (*model.Client, error)
	GetClient(ctx context.Context, clientID string) (*model.Client, error)

	CreateUser(ctx context.Context, loginID string) (*model.User, error)
	GetUser(ctx context.Context, loginID string) (*model.User, error)
}

var (
	// ErrNotFound はデータが見つからなかったときのエラー。
	ErrNotFound = errors.New("repository.ErrNotFound")

	// ErrNotSaved はデータが保存できなかったときのエラー。
	ErrNotSaved = errors.New("repository.ErrNotSaved")
)
