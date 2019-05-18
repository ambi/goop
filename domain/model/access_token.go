package model

import (
	"log"
	"time"

	"github.com/ambi/goop/domain/oidc"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/jwt"
)

// AccessToken is a type for access tokens.
type AccessToken struct {
	jwt.Claims
	Subtype string   `json:"subtype,omitempty"`
	Scopes  []string `json:"scopes,omitempty"`
	OP      *OP      `json:"-"`
	User    *User    `json:"-"`
}

// NewAccessToken creates a new access token.
func NewAccessToken(now time.Time, op *OP, resourceID string, user *User, expiresIn int) *AccessToken {
	expiry := now.Add(time.Duration(expiresIn) * time.Second)

	accessToken := &AccessToken{
		Claims: jwt.Claims{
			Issuer:    op.Issuer,
			Subject:   user.UUID,
			Audience:  []string{resourceID},
			Expiry:    jwt.NewNumericDate(expiry),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        uuid.New().String(),
		},
		Subtype: oidc.SubtypeAccessToken,
		OP:      op,
		User:    user,
	}

	return accessToken
}

// TokenType returns the type of the access token.
func (t *AccessToken) TokenType() string {
	return oidc.TokenTypeBearer
}

// String returns string representation for the access token.
func (t *AccessToken) String() string {
	raw, err := jwt.Signed(t.OP.JWTSigner).Claims(t).CompactSerialize()
	if err != nil {
		log.Fatalf("IDToken.String error=%v\n", err)
	}
	return raw
}
