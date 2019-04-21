package model

import (
	"log"
	"time"

	"github.com/ambi/goop/domain/oidc"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2/jwt"
)

// AccessToken はアクセストークンの型。
type AccessToken struct {
	jwt.Claims
	Subtype string   `json:"subtype,omitempty"`
	Scopes  []string `json:"scopes,omitempty"`
	OP      *OP      `json:"-"`
}

// NewAccessToken は新しいアクセストークンを生成する。
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
	}

	return accessToken
}

// TokenType はアクセストークンタイプ (bearer, mac, ...) を返す。
func (t *AccessToken) TokenType() string {
	return oidc.TokenTypeBearer
}

// String は アクセストークンの文字列表現を返す。
func (t *AccessToken) String() string {
	raw, err := jwt.Signed(t.OP.JWTSigner).Claims(t).CompactSerialize()
	if err != nil {
		log.Fatalf("IDToken.String error=%v\n", err)
	}
	return raw
}
