package usecase

import "errors"

var (
	// ErrUserUnauthenticated is an error when a user cannot be authenticated.
	ErrUserUnauthenticated = errors.New("usecase.ErrUserUnauthenticated")
)
