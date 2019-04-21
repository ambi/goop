package model

// AuthorizationCode は OIDC 認可コードを表した型。
type AuthorizationCode struct {
	UUID  string
	Code  string
	Nonce string
	User  *User
}
