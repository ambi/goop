package model

import "encoding/json"

// TokenResponse is a type for responses of the OAuth 2.0 token endpoint.
type TokenResponse struct {
	AccessToken *AccessToken
	IDToken     *IDToken
	TokenType   string
	ExpiresIn   int
}

// MarshalJSON converts a TokenResponse to JSON.
func (res *TokenResponse) MarshalJSON() ([]byte, error) {
	obj := map[string]interface{}{
		"access_token": res.AccessToken.String(),
		"id_token":     res.IDToken.String(),
		"token_type":   res.TokenType,
		"expires_in":   res.ExpiresIn,
	}
	return json.Marshal(obj)
}
