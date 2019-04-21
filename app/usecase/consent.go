package usecase

import (
	"github.com/ambi/goop/domain/db"
)

// Consent は同意処理のユースケース。
type Consent struct {
	dao db.DAO
}

// NewConsent は Consent ユースケースを生成する。
func NewConsent(dao db.DAO) *Consent {
	return &Consent{dao}
}

// Call は同意処理を行う。
func (uc *Consent) Call() {
}
