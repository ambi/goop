package model

import "net/url"

// AuthorizeError is a type for authorization errors.
type AuthorizeError struct {
	StatusCode int
	Message    string
	State      string
}

func (err *AuthorizeError) Error() string {
	return err.Message
}

// ToQuery converts an error to a URL query string.
func (err *AuthorizeError) ToQuery() string {
	values := url.Values{}
	values.Add("error", err.Message)
	if err.State != "" {
		values.Add("state", err.State)
	}

	return values.Encode()
}
