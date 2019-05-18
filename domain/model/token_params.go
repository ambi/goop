package model

import (
	"net/http"
)

// TokenParams is a type for parameters for the OAuth 2.0 token endpoint.
type TokenParams struct {
	GrantType    string
	Code         string
	RedirectURI  string
	ClientID     string
	ClientSecret string
}

// Valid validates a TokenParams.
func (params *TokenParams) Valid(client *Client, authzCode *AuthorizationCode) *ClientError {
	if client == nil {
		err := &ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_client",
		}
		return err
	}

	if client.ClientSecret != params.ClientSecret {
		err := &ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_client",
		}
		return err
	}

	if !client.IsValidRedirectURI(params.RedirectURI) {
		err := &ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_request",
		}
		return err
	}

	if params.GrantType == "" {
		err := &ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_request",
		}
		return err
	}
	if params.GrantType != "authorization_code" {
		err := &ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    "unsupported_grant_type",
		}
		return err
	}

	if authzCode == nil {
		err := &ClientError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_grant",
		}
		return err
	}

	return nil
}
