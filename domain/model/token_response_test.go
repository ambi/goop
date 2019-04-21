package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenResponse_MarshalJSON(t *testing.T) {
	now := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	op := NewTestOP()
	user := &User{UUID: "0093803a-ed45-4d2e-8f69-d34d7228e5da"}
	at := NewAccessToken(now, op, "resource", user, 3600)
	it := NewIDToken(now, op, "client1", user, 3600, "nonce1")
	res := TokenResponse{
		AccessToken: at,
		IDToken:     it,
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}

	j, err := res.MarshalJSON()

	assert.NoError(t, err)
	assert.Contains(t, string(j), `"access_token":"eyJ`)
	assert.Contains(t, string(j), `"id_token":"eyJ`)
}
