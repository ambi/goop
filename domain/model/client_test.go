package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_IsValidRedirectURI(t *testing.T) {
	client := &Client{
		RedirectURIs: []string{"http://example.com", "http://example.jp"},
	}
	testCases := []struct {
		redirectURI string
		want        bool
	}{
		{"http://example.com", true},
		{"http://example.jp", true},
		{"http://invalid.com", false},
		{"http://example.com/index.html", false},
	}
	for _, tc := range testCases {
		got := client.IsValidRedirectURI(tc.redirectURI)

		assert.Equal(t, tc.want, got)
	}
}
