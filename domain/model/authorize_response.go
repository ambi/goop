package model

import "net/url"

// AuthorizeResponse is a type for responses of the token endpoint.
type AuthorizeResponse struct {
	Code  string
	State string
}

// ToQuery converts an AuthorizeResponse to a URL query string.
func (resp *AuthorizeResponse) ToQuery() string {
	values := url.Values{}
	values.Add("code", resp.Code)
	if resp.State != "" {
		values.Add("state", resp.State)
	}

	return values.Encode()
}
