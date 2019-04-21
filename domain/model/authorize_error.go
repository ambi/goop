package model

import "net/url"

// AuthorizeError は認可エラーの型。
type AuthorizeError struct {
	StatusCode int
	Message    string
	State      string
}

func (err *AuthorizeError) Error() string {
	return err.Message
}

// ToQuery はエラーを URL クエリ文字列に変換する。
func (err *AuthorizeError) ToQuery() string {
	values := url.Values{}
	values.Add("error", err.Message)
	if err.State != "" {
		values.Add("state", err.State)
	}

	return values.Encode()
}
