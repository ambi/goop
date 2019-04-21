package model

import (
	"net/http"
	"testing"

	"github.com/ambi/goop/domain/oidc"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizeParams_Valid(t *testing.T) {
	scope := "openid"
	clientID := "client-id1"
	redirectURI := "http://localhost/callback"
	state := "state1"
	responseMode := ""
	nonce := "nonce1"
	client := &Client{ClientID: clientID, RedirectURIs: []string{redirectURI}}
	testCases := []struct {
		params AuthorizeParams
		client *Client
		want   *AuthorizeError
	}{
		{AuthorizeParams{scope, oidc.ResponseTypeCode, clientID, redirectURI, state, responseMode, nonce}, client, nil},
		{AuthorizeParams{scope, oidc.ResponseTypeCode, clientID, redirectURI, state, responseMode, nonce}, nil, &AuthorizeError{StatusCode: http.StatusUnauthorized, Message: "invalid_client"}},
		{AuthorizeParams{scope, oidc.ResponseTypeCode, clientID, "invalid", state, responseMode, nonce}, client, &AuthorizeError{StatusCode: http.StatusBadRequest, Message: "invalid_request"}},
		{AuthorizeParams{scope, "", clientID, redirectURI, state, responseMode, nonce}, client, &AuthorizeError{StatusCode: http.StatusFound, Message: "invalid_request", State: state}},
		{AuthorizeParams{scope, "invalid", clientID, redirectURI, state, responseMode, nonce}, client, &AuthorizeError{StatusCode: http.StatusFound, Message: "unsupported_response_type", State: state}},
		{AuthorizeParams{"invalid", oidc.ResponseTypeCode, clientID, redirectURI, state, responseMode, nonce}, client, &AuthorizeError{StatusCode: http.StatusFound, Message: "invalid_scope", State: state}},
		{AuthorizeParams{scope, oidc.ResponseTypeCode, clientID, redirectURI, state, "invalid", nonce}, client, &AuthorizeError{StatusCode: http.StatusFound, Message: "unsupported_response_mode", State: state}},
	}
	for _, tc := range testCases {
		got := tc.params.Valid(tc.client)

		assert.Equal(t, tc.want, got)
	}
}
