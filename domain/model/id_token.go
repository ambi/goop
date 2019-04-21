package model

import (
	"log"
	"time"

	"github.com/ambi/goop/domain/oidc"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/jwt"
)

// IDToken は ID トークン型。
type IDToken struct {
	jwt.Claims
	Subtype  string           `json:"subtype,omitempty"`
	AuthTime *jwt.NumericDate `json:"auth_time,omitempty"`
	Nonce    string           `json:"nonce,omitempty"`
	// Acr, Amr, Azp
	OP *OP `json:"-"`
}

// NewIDToken は新しい ID トークンを生成する。
func NewIDToken(now time.Time, op *OP, clientID string, user *User, expiresIn int, nonce string) *IDToken {
	expiry := now.Add(time.Duration(expiresIn) * time.Second)

	idToken := &IDToken{
		Claims: jwt.Claims{
			Issuer:    op.Issuer,
			Subject:   user.UUID,
			Audience:  []string{clientID},
			Expiry:    jwt.NewNumericDate(expiry),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        uuid.New().String(),
		},
		Subtype:  oidc.SubtypeIDToken,
		AuthTime: jwt.NewNumericDate(now),
		Nonce:    nonce,
		OP:       op,
	}

	return idToken
}

// String は ID トークンの文字列表現を返す。
func (t *IDToken) String() string {
	raw, err := jwt.Signed(t.OP.JWTSigner).Claims(t).CompactSerialize()
	if err != nil {
		log.Fatalf("IDToken.String error=%v\n", err)
	}
	return raw
}
