package model

import (
	"net/http"
)

// RevokeParams is a type for parameters of the rovocation endpoint.
type RevokeParams struct {
	Token         string
	TokenTypeHint string
	ClientID      string
	ClientSecret  string
}

// Valid validates RevokeParams.
func (params *RevokeParams) Valid(client *Client) *ClientError {
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

	return nil
}
