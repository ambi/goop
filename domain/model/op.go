package model

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/url"
	"path"

	"github.com/ambi/goop/domain/oidc"
	"gopkg.in/square/go-jose.v2"
)

// OP は OpenID Provider の型。
type OP struct {
	Issuer                           string   `json:"issuer"`
	AuthorizationEndpoint            string   `json:"authorization_endpoint"`
	TokenEndpoint                    string   `json:"token_endpoint"`
	JwksURI                          string   `json:"jwks_uri"`
	ScopesSupported                  []string `json:"scopes_supported"`
	ResponseTypeSupported            []string `json:"response_types_supported"`
	ResponseModesSupported           []string `json:"response_modes_supported"`
	GrantTypesSupported              []string `json:"grant_types_supported"`
	SubjectTypesSupported            []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported"`
	ClaimsSupported                  []string `json:"claims_supported"`

	PrivateJWKS jose.JSONWebKeySet `json:"-"`
	PublicJWKS  jose.JSONWebKeySet `json:"-"`
	JWTSigner   jose.Signer        `json:"-"`
}

// NewSimpleOP はシンプルな OP を生成する。
func NewSimpleOP(issuer string, privateJWKS jose.JSONWebKeySet) *OP {
	url, err := url.Parse(issuer)
	if err != nil {
		log.Printf("NewOP Error: err=%v", err)
		return nil
	}
	basePath := url.Path

	url.Path = path.Join(basePath, "authorize")
	authorizationEndpoint := url.String()

	url.Path = path.Join(basePath, "token")
	tokenEndpoint := url.String()

	url.Path = path.Join(basePath, "jwks")
	jwksURI := url.String()

	publicJWKS := jose.JSONWebKeySet{Keys: []jose.JSONWebKey{}}
	for _, jwk := range privateJWKS.Keys {
		publicJWKS.Keys = append(publicJWKS.Keys, jwk.Public())
	}

	var jwtSigner jose.Signer
	if len(privateJWKS.Keys) >= 1 {
		jwtSigner, err = jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: privateJWKS.Keys[0]}, (&jose.SignerOptions{}).WithType("JWT"))
		if err != nil {
			log.Printf("oidc.jwks init error: err=%v", err)
		}
	}

	op := &OP{
		Issuer:                           issuer,
		AuthorizationEndpoint:            authorizationEndpoint,
		TokenEndpoint:                    tokenEndpoint,
		JwksURI:                          jwksURI,
		ScopesSupported:                  oidc.ScopesSupported,
		ResponseTypeSupported:            oidc.ResponseTypesSupported,
		ResponseModesSupported:           oidc.ResponseModesSupported,
		GrantTypesSupported:              oidc.GrantTypesSupported,
		SubjectTypesSupported:            oidc.SubjectTypesSupported,
		IDTokenSigningAlgValuesSupported: oidc.IDTokenSigningAlgValuesSupported,
		ClaimsSupported:                  oidc.ClaimsSupported,
		PrivateJWKS:                      privateJWKS,
		PublicJWKS:                       publicJWKS,
		JWTSigner:                        jwtSigner,
	}
	return op
}

// NewJSONWebKeySetFromPEMs は PEM の配列から JWK Set を生成して返す。
func NewJSONWebKeySetFromPEMs(pems []string) jose.JSONWebKeySet {
	const (
		keyAlg = string(jose.RS256)
		keyUse = "sig"
	)

	jwks := make([]jose.JSONWebKey, 0, len(pems))

	for _, p := range pems {
		block, _ := pem.Decode([]byte(p))
		if block == nil {
			log.Fatalf("NewJSONWebKeySetFromPEMs error")
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatalf("NewJSONWebKeySetFromPEMs error: err=%v", err)
		}

		publicKeyDER := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
		fingerprint := fmt.Sprintf("%x", sha1.Sum(publicKeyDER))

		jwk := jose.JSONWebKey{
			Key:       privateKey,
			KeyID:     fingerprint,
			Algorithm: keyAlg,
			Use:       keyUse,
		}

		jwks = append(jwks, jwk)
	}

	return jose.JSONWebKeySet{Keys: jwks}
}
