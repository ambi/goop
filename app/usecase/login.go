package usecase

import (
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/gorilla/sessions"
)

// Login is a use case for login.
type Login struct {
	repo db.Repository
}

// NewLogin creates a Login use case.
func NewLogin(repo db.Repository) *Login {
	return &Login{repo}
}

// Call processes a login.
func (uc *Login) Call(sess *sessions.Session, loginID, password string) (*model.User, error) {
	// TODO
	if loginID != "test1@example.com" || password != "hogehoge" {
		return nil, ErrUserUnauthenticated
	}

	user := &model.User{
		LoginID: loginID,
	}
	sess.Values["current_user"] = user.LoginID
	return user, nil
}
