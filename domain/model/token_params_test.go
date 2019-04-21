package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenParams_Valid(t *testing.T) {
	redirectURI := "http://localhost/callback"
	client := &Client{ClientID: "client-id1", ClientSecret: "client-secret1", RedirectURIs: []string{redirectURI}}
	authzCode := &AuthorizationCode{Code: "code1"}
	testCases := []struct {
		params    TokenParams
		client    *Client
		authzCode *AuthorizationCode
		want      *TokenError
	}{
		{TokenParams{"authorization_code", authzCode.Code, redirectURI, client.ClientID, client.ClientSecret}, client, authzCode, nil},
		{TokenParams{"authorization_code", authzCode.Code, redirectURI, "invalid", client.ClientSecret}, nil, authzCode, &TokenError{StatusCode: http.StatusBadRequest, Message: "invalid_client"}},
		{TokenParams{"authorization_code", authzCode.Code, redirectURI, client.ClientID, "invalid"}, client, authzCode, &TokenError{StatusCode: http.StatusBadRequest, Message: "invalid_client"}},
		{TokenParams{"authorization_code", authzCode.Code, "invalid", client.ClientID, client.ClientSecret}, client, authzCode, &TokenError{StatusCode: http.StatusBadRequest, Message: "invalid_request"}},
		{TokenParams{"", authzCode.Code, redirectURI, client.ClientID, client.ClientSecret}, client, authzCode, &TokenError{StatusCode: http.StatusBadRequest, Message: "invalid_request"}},
		{TokenParams{"invalid", authzCode.Code, redirectURI, client.ClientID, client.ClientSecret}, client, authzCode, &TokenError{StatusCode: http.StatusBadRequest, Message: "unsupported_grant_type"}},
		{TokenParams{"authorization_code", "invalid", redirectURI, client.ClientID, client.ClientSecret}, client, nil, &TokenError{StatusCode: http.StatusBadRequest, Message: "invalid_grant"}},
	}
	for _, tc := range testCases {
		got := tc.params.Valid(tc.client, tc.authzCode)

		assert.Equal(t, tc.want, got)
	}
}
