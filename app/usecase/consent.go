package usecase

import (
	"github.com/ambi/goop/domain/db"
)

// Consent is an use case for consent.
type Consent struct {
	repo db.Repository
}

// NewConsent creates a Consent use case.
func NewConsent(repo db.Repository) *Consent {
	return &Consent{repo}
}

// Call processes consent.
func (uc *Consent) Call() {
	// TODO
}
