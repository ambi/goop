package model

import (
	"net/http"
	"strings"

	"github.com/ambi/goop/domain/oidc"
)

// AuthorizeParams は Authorization Endpoint のパラメータをひとまとめにした型。
type AuthorizeParams struct {
	Scope        string
	ResponseType string
	ClientID     string
	RedirectURI  string
	State        string
	ResponseMode string
	Nonce        string
}

// Valid は AuthorizeParams のバリデーションを行う。
func (params *AuthorizeParams) Valid(client *Client) *AuthorizeError {
	if client == nil {
		err := &AuthorizeError{
			StatusCode: http.StatusUnauthorized,
			Message:    "invalid_client",
		}
		return err
	}
	if !client.IsValidRedirectURI(params.RedirectURI) {
		err := &AuthorizeError{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid_request",
		}
		return err
	}

	if params.ResponseType == "" {
		err := &AuthorizeError{
			StatusCode: http.StatusFound,
			Message:    "invalid_request",
			State:      params.State,
		}
		return err
	}
	for _, responseType := range params.ResponseTypes() {
		if !oidc.IsSupportedResponseType(responseType) {
			err := &AuthorizeError{
				StatusCode: http.StatusFound,
				Message:    "unsupported_response_type",
				State:      params.State,
			}
			return err
		}
	}

	for _, scope := range params.Scopes() {
		if !oidc.IsSupportedScope(scope) {
			err := &AuthorizeError{
				StatusCode: http.StatusFound,
				Message:    "invalid_scope",
				State:      params.State,
			}
			return err
		}
	}

	if params.ResponseMode != "" && !oidc.IsSupportedResponseMode(params.ResponseMode) {
		err := &AuthorizeError{
			StatusCode: http.StatusFound,
			Message:    "unsupported_response_mode",
			State:      params.State,
		}
		return err
	}

	return nil
}

// ResponseTypes は response_type を配列にして返す。
func (params *AuthorizeParams) ResponseTypes() []string {
	return strings.Split(params.ResponseType, " ")
}

// Scopes は scope を配列にして返す。
func (params *AuthorizeParams) Scopes() []string {
	return strings.Split(params.Scope, " ")
}
