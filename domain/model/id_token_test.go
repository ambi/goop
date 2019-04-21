package model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewIDToken(t *testing.T) {
	now := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	op := NewTestOP()
	user := &User{UUID: "0093803a-ed45-4d2e-8f69-d34d7228e5da"}

	testCases := []struct {
		op        *OP
		clientID  string
		user      *User
		expiresIn int
		nonce     string
	}{
		{op, "client1", user, 3600, "nonce1"},
		{op, "client2", user, 1, "nonce2"},
	}
	for _, tc := range testCases {
		nowInt := now.Unix()
		expiryInt := nowInt + int64(tc.expiresIn)

		got := NewIDToken(now, tc.op, tc.clientID, tc.user, tc.expiresIn, tc.nonce)

		assert.Equal(t, tc.op.Issuer, got.Claims.Issuer)
		assert.Equal(t, tc.user.UUID, got.Claims.Subject)
		assert.Equal(t, []string{tc.clientID}, []string(got.Claims.Audience))
		assert.Equal(t, expiryInt, int64(*got.Claims.Expiry))
		assert.Equal(t, nowInt, int64(*got.Claims.NotBefore))
		assert.Equal(t, nowInt, int64(*got.Claims.IssuedAt))
		assert.NotEmpty(t, got.Claims.ID)
	}
}

func TestIDToken_String(t *testing.T) {
	now := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	op := NewTestOP()
	user := &User{UUID: "0093803a-ed45-4d2e-8f69-d34d7228e5da"}
	token := NewIDToken(now, op, "resource", user, 3600, "nonce")
	jwted := token.String()

	assert.True(t, strings.HasPrefix(jwted, "eyJ"))
	assert.Equal(t, 2, strings.Count(jwted, "."))
}
