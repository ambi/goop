package usecase

import "errors"

var (
	// ErrUserUnauthenticated はユーザ認証に失敗したときのエラー。
	ErrUserUnauthenticated = errors.New("usecase.ErrUserUnauthenticated")
)
