package model

import (
	"net/http"
)

// TokenParams は Token Endpoint のパラメータをまとめた型。
type TokenParams struct {
	GrantType    string
	Code         string
	RedirectURI  string
	ClientID     string
	ClientSecret string
}

// Valid は TokenParams のバリデーションを行う。
func (params *TokenParams) Valid(client *Client, authzCode *AuthorizationCode) *TokenError {
	if client == nil {
		err := &TokenError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_client",
		}
		return err
	}

	if client.ClientSecret != params.ClientSecret {
		err := &TokenError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_client",
		}
		return err
	}

	if !client.IsValidRedirectURI(params.RedirectURI) {
		err := &TokenError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_request",
		}
		return err
	}

	if params.GrantType == "" {
		err := &TokenError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_request",
		}
		return err
	}
	if params.GrantType != "authorization_code" {
		err := &TokenError{
			StatusCode: http.StatusBadRequest,
			Message:    "unsupported_grant_type",
		}
		return err
	}

	if authzCode == nil {
		err := &TokenError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_grant",
		}
		return err
	}

	return nil
}
