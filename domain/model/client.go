package model

// Client は OIDC Client を表した型。
type Client struct {
	UUID         string
	ClientID     string
	ClientSecret string
	Name         string
	RedirectURIs []string
}

// IsValidRedirectURI は登録ずみのリダイレクト URI かどうかを返す。
func (c *Client) IsValidRedirectURI(redirectURI string) bool {
	for _, uri := range c.RedirectURIs {
		if uri == redirectURI {
			return true
		}
	}
	return false
}
