package model

import "net/url"

// AuthorizeResponse は Token Endpoint のレスポンスを抽象化した型。
type AuthorizeResponse struct {
	Code  string
	State string
}

// ToQuery は AuthorizeResponse を URL クエリ文字列に変換する。
func (resp *AuthorizeResponse) ToQuery() string {
	values := url.Values{}
	values.Add("code", resp.Code)
	if resp.State != "" {
		values.Add("state", resp.State)
	}

	return values.Encode()
}
