package usecase

import (
	"github.com/ambi/goop/domain/db"
	"github.com/ambi/goop/domain/model"
	"github.com/gorilla/sessions"
)

// Login はログイン処理のユースケース。
type Login struct {
	dao db.DAO
}

// NewLogin は Login ユースケースを生成する。
func NewLogin(dao db.DAO) *Login {
	return &Login{dao}
}

// Call はログイン処理を行う。
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
