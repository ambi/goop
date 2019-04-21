package model

import (
	"strings"
	"testing"
	"time"

	"github.com/ambi/goop/domain/oidc"
	"github.com/stretchr/testify/assert"
)

func TestNewAccessToken(t *testing.T) {
	now := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	op := NewTestOP()
	user := &User{UUID: "0093803a-ed45-4d2e-8f69-d34d7228e5da"}

	testCases := []struct {
		op         *OP
		resourceID string
		user       *User
		expiresIn  int
	}{
		{op, "resource1", user, 3600},
		{op, "resource2", user, 1},
	}
	for _, tc := range testCases {
		nowInt := now.Unix()
		expiryInt := nowInt + int64(tc.expiresIn)

		got := NewAccessToken(now, tc.op, tc.resourceID, tc.user, tc.expiresIn)

		assert.Equal(t, tc.op.Issuer, got.Claims.Issuer)
		assert.Equal(t, tc.user.UUID, got.Claims.Subject)
		assert.Equal(t, []string{tc.resourceID}, []string(got.Claims.Audience))
		assert.Equal(t, expiryInt, int64(*got.Claims.Expiry))
		assert.Equal(t, nowInt, int64(*got.Claims.NotBefore))
		assert.Equal(t, nowInt, int64(*got.Claims.IssuedAt))
		assert.NotEmpty(t, got.Claims.ID)
	}
}

func TestAccessToken_TokenType(t *testing.T) {
	now := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	op := NewTestOP()
	user := &User{UUID: "0093803a-ed45-4d2e-8f69-d34d7228e5da"}
	token := NewAccessToken(now, op, "resource", user, 3600)

	assert.Equal(t, oidc.TokenTypeBearer, token.TokenType())
}

func TestAccessToken_String(t *testing.T) {
	now := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	op := NewTestOP()
	user := &User{UUID: "0093803a-ed45-4d2e-8f69-d34d7228e5da"}
	token := NewAccessToken(now, op, "resource", user, 3600)
	jwted := token.String()

	assert.True(t, strings.HasPrefix(jwted, "eyJ"))
	assert.Equal(t, 2, strings.Count(jwted, "."))
}
