package oidc

import "gopkg.in/square/go-jose.v2"

const (
	// ClaimSub は sub クレームを表現する。
	ClaimSub = "sub"
	// ClaimIss は iss クレームを表現する。
	ClaimIss = "iss"
	// ClaimAud は aud クレームを表現する。
	ClaimAud = "aud"
	// ClaimExp は exp クレームを表現する。
	ClaimExp = "exp"
	// ClaimIat は iat クレームを表現する。
	ClaimIat = "iat"

	// GrantTypeAuthorizationCode は認可コードグラント grant_type=authorization_code を表現する。
	GrantTypeAuthorizationCode = "authorization_code"
	// GrantTypeImplicit はインプリシットグラント grant_type=implicit を表現する。
	GrantTypeImplicit = "implicit"

	// ResponseModeQuery は response_type=query (?state=...) を表現する。
	ResponseModeQuery = "query"
	// ResponseModeFragment は response_type=fragment (#state=...) を表現する。
	ResponseModeFragment = "fragment"
	// ResponseModeFormPost は response_type=form_post (<form>...</form>) を表現する。
	ResponseModeFormPost = "form_post"

	// ResponseTypeCode は Authorization Code Flow の response_type=code を表現する。
	ResponseTypeCode = "code"
	// ResponseTypeToken は Implicit Flow の response_type=token を表現する。
	ResponseTypeToken = "token"
	// ResponseTypeIDToken は Implicit Flow の response_type=id_token を表現する。
	ResponseTypeIDToken = "id_token"

	// ScopeOpenID は openid スコープを表現する。
	ScopeOpenID = "openid"
	// ScopeProfile は profile スコープを表現する。
	ScopeProfile = "profile"

	// SubjectTypePairwise は subject_type=pairwise を表現する。
	SubjectTypePairwise = "pairwise"
	// SubjectTypePublic は subject_type=Public を表現する。
	SubjectTypePublic = "public"

	// SubtypeAccessToken はトークンの種類がアクセストークンであることを表現する。
	SubtypeAccessToken = "access_token"
	// SubtypeIDToken はトークンの種類が ID トークンであることを表現する。
	SubtypeIDToken = "id_token"

	// TokenTYpeBearer はアクセストークンタイプが Bearer であることを表現する。
	TokenTypeBearer = "bearer"
)

// ClaimsSupported はサポートするクレームのスライス。
var ClaimsSupported = []string{
	ClaimIss,
	ClaimSub,
	ClaimAud,
	ClaimExp,
	ClaimIat,
}

// GrantTypesSupported はサポートする grant_type のスライス。
var GrantTypesSupported = []string{
	GrantTypeAuthorizationCode,
	GrantTypeImplicit,
}

// IDTokenSigningAlgValuesSupported はサポートする ID トークン署名アルゴリズムのスライス。
var IDTokenSigningAlgValuesSupported = []string{
	string(jose.RS256),
}

// ResponseModesSupported はサポートする response_mode のスライス。
var ResponseModesSupported = []string{
	ResponseModeQuery,
	ResponseModeFragment,
	ResponseModeFormPost,
}

// ResponseTypesSupported はサポートする response_type のスライス。
var ResponseTypesSupported = []string{
	ResponseTypeCode,
	ResponseTypeToken,
	ResponseTypeIDToken,
}

// ScopesSupported はサポートする scope のスライス。
var ScopesSupported = []string{
	ScopeOpenID,
	ScopeProfile,
}

// SubjectTypesSupported はサポートする subject_type のスライス。
var SubjectTypesSupported = []string{
	SubjectTypePublic,
}
