package model

// TokenError は認可エラーの型。
type TokenError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"error"`
}

func (err *TokenError) Error() string {
	return err.Message
}
