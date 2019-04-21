package model

import "encoding/json"

// TokenResponse は Token Endpoint の成功レスポンスを抽象化した型。
type TokenResponse struct {
	AccessToken *AccessToken
	IDToken     *IDToken
	TokenType   string
	ExpiresIn   int
}

// MarshalJSON は Token Response を JSON 形式に変換する。
func (res *TokenResponse) MarshalJSON() ([]byte, error) {
	obj := map[string]interface{}{
		"access_token": res.AccessToken.String(),
		"id_token":     res.IDToken.String(),
		"token_type":   res.TokenType,
		"expires_in":   res.ExpiresIn,
	}
	return json.Marshal(obj)
}
