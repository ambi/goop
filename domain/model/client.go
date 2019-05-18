package model

// Client is a type for OAuth 2.0 clients.
type Client struct {
	UUID         string
	ClientID     string
	ClientSecret string
	Name         string
	RedirectURIs []string
}

// IsValidRedirectURI validates a redirect URI.
func (c *Client) IsValidRedirectURI(redirectURI string) bool {
	for _, uri := range c.RedirectURIs {
		if uri == redirectURI {
			return true
		}
	}
	return false
}
