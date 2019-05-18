package oidc

import "gopkg.in/square/go-jose.v2"

const (
	// ClaimSub represents sub claim.
	ClaimSub = "sub"
	// ClaimIss represents iss claim.
	ClaimIss = "iss"
	// ClaimAud represents aud claim.
	ClaimAud = "aud"
	// ClaimExp represents exp claim.
	ClaimExp = "exp"
	// ClaimIat represents iat claim.
	ClaimIat = "iat"

	// GrantTypeAuthorizationCode represents grant_type=authorization_code.
	GrantTypeAuthorizationCode = "authorization_code"
	// GrantTypeImplicit represents grant_type=implicit.
	GrantTypeImplicit = "implicit"

	// ResponseModeQuery represents response_type=query (?state=...).
	ResponseModeQuery = "query"
	// ResponseModeFragment represents response_type=fragment (#state=...).
	ResponseModeFragment = "fragment"
	// ResponseModeFormPost represents response_type=form_post (<form>...</form>).
	ResponseModeFormPost = "form_post"

	// ResponseTypeCode represents response_type=code.
	ResponseTypeCode = "code"
	// ResponseTypeToken represents response_type=token.
	ResponseTypeToken = "token"
	// ResponseTypeIDToken represents response_type=id_token.
	ResponseTypeIDToken = "id_token"

	// ScopeOpenID represents openid scope.
	ScopeOpenID = "openid"
	// ScopeProfile represents profile scope.
	ScopeProfile = "profile"

	// SubjectTypePairwise represents subject_type=pairwise.
	SubjectTypePairwise = "pairwise"
	// SubjectTypePublic represents subject_type=Public.
	SubjectTypePublic = "public"

	// SubtypeAccessToken represents an access token.
	SubtypeAccessToken = "access_token"
	// SubtypeIDToken represents an ID token.
	SubtypeIDToken = "id_token"

	// TokenTypeBearer represents a bearer token.
	TokenTypeBearer = "bearer"
)

// ClaimsSupported is a slice of supported claims.
var ClaimsSupported = []string{
	ClaimIss,
	ClaimSub,
	ClaimAud,
	ClaimExp,
	ClaimIat,
}

// GrantTypesSupported is a slice of supported grant types.
var GrantTypesSupported = []string{
	GrantTypeAuthorizationCode,
	GrantTypeImplicit,
}

// IDTokenSigningAlgValuesSupported is a slice of supported ID token signing algorithms.
var IDTokenSigningAlgValuesSupported = []string{
	string(jose.RS256),
}

// ResponseModesSupported is a slice of supported response modes.
var ResponseModesSupported = []string{
	ResponseModeQuery,
	ResponseModeFragment,
	ResponseModeFormPost,
}

// ResponseTypesSupported is a slice of supported response types.
var ResponseTypesSupported = []string{
	ResponseTypeCode,
	ResponseTypeToken,
	ResponseTypeIDToken,
}

// ScopesSupported is a slice of supported scopes.
var ScopesSupported = []string{
	ScopeOpenID,
	ScopeProfile,
}

// SubjectTypesSupported is a slice of supported subject types.
var SubjectTypesSupported = []string{
	SubjectTypePublic,
}
