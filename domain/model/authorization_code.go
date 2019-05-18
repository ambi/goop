package model

// AuthorizationCode is a type for OIDC authorization codes.
type AuthorizationCode struct {
	UUID  string
	Code  string
	Nonce string
	User  *User
}
