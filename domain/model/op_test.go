package model

import (
	"testing"

	"github.com/ambi/goop/domain/oidc"
	"github.com/stretchr/testify/assert"
	"gopkg.in/square/go-jose.v2"
)

const privateKeyPEM = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwDbhuR4WQw45kdhhdHhBGOItNcDaNVA5H96EeYuFA18owytt
Ve8N/27xbcOAVsE/cytYIzS9BZ0TajAMfdxqNGFaGs9hF2DqSOCQCZEljcBjzyvm
JHArIsJAeAmsACkTU9iczJaCfsuI8nU07ZGoejRWc+n3T7+Ezv7J8ByMQ7gAbsee
916xl7yeXqRc9sh8Ashhh4vG0qUpZUgE/3U5XVc7VgSnuIayez6MUkUW+W7uoNvo
aK66X5bpFlvLqdn9nw6yQqT4MUvs+8R0IqrjpCWo/M5qYjeiIqBGZ+speHLIB2Bl
U6/hmGHveTqLRYq8NFdgdawHkPAg3Xnnx4X8rwIDAQABAoIBAAdKjaluAL4ubfSg
VqIKZb3B5qEhXaWkE2aoFDJgHn2QQIWWwiD+XnOQC8x0HAwubG+79y3utDTgTno6
tF6Z29V80+QsdaxjWr/E9PHH9bq4z5BuQ96h+qDo6XUkOMgQxtDpaJlCJFUW92zr
EUTpv5BTJ+xhtaXqx+jqNVLMYm23HkinE185qN4xD2vcMlC/RFNPiTsKxeMBb8/H
gyLuxsd/FAW8uuA56v7Adq7tdbeesRttK93xonE6738uaivrvSbwUZ+htUCBGVDv
Uo7sn5cPl/JsyGpCGLvLNCmVMfDXbMJkfsooZysNuZTK+frgFWK9QSlXZaK9Lgje
IFmhofECgYEA367/FgYg+H7Wodd4LKwP52s5ACoMYCJKP/HTD4RXbNWwzgYPFiVf
9A67Jcn9xt2FCdhtKk4qX4zbq/kuRezcWkRBrsSiu3ZiJeZqsBkzrf8eqSR56NSQ
sgyqdjq1cbsdr8a55fJc7FHCkYOIGzBn7kHrO+VIM1kIA8pob+/nWdUCgYEA2/v8
2DQR2GrgDbu65W8DBmy8bamzNKo/BijD82AlPjAideDfd/ApFuozV8QyPUGPnHsr
4WmniHSsv0OzQLWV7t6qWwcJWscHP1tPG79oHzptddq59F466olM75tBTNiVZK4W
ntLjMeHRwp4cAyYEgKE1/MDv+4qdr0vlEBEVGnMCgYEAt8xYqGUy1F2MyYF+qt5b
VyHgvYTPlHK4piBz+E0bdT+Fv1R1MxJU7PrpxpxaXf0J41QmJ1wBL3BT1NS0tBpz
2ew6OHw052rYLSJPShH7SB7Yt81UKv/3QYZszydgjmpdc2EMwajLuBdalC0LOqQk
0j7yJs447JBcqJi6BjiP1l0CgYEAwasSCcGlqFNHF7AbtVvLXP+j0CPjqkzrmV+Q
S5Zsk1hLLl4gPvZHJm5fbzhCRp2OcmQZ2KRsovydDPHsQN7cteSANA141dt87PFV
LwePAlctAOHkblf7JHpmKlgT4DSZKX7+WSsua770LZOG89qnghrrba7qWBjMBAUc
D84KNzMCgYBZWCYEq8gTXDaq33bewYO+Bo5uewXGrj/9U5JkL6fx0UqtB1Fkr+xg
YObbtdbLq6UITyuudmaHeLeOeYC3FbVQuJ6+iluJFDMJqmVz0inHcdzCURJAsAvO
Y1BDOqcjK4iU90b1bzFs55Pngu9c1lFEtqchQjADvggTEi0/7aqcLg==
-----END RSA PRIVATE KEY-----`

func TestNewSimpleOP(t *testing.T) {
	privateJWKS := NewJSONWebKeySetFromPEMs([]string{privateKeyPEM})
	publicJWKS := jose.JSONWebKeySet{Keys: []jose.JSONWebKey{privateJWKS.Keys[0].Public()}}
	jwtSigner, _ := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: privateJWKS.Keys[0]}, (&jose.SignerOptions{}).WithType("JWT"))

	got := NewTestOP()

	assert.Equal(t, "http://example.com/op", got.Issuer)
	assert.Equal(t, "http://example.com/op/authorize", got.AuthorizationEndpoint)
	assert.Equal(t, "http://example.com/op/token", got.TokenEndpoint)
	assert.Equal(t, "http://example.com/op/jwks", got.JwksURI)
	assert.Equal(t, oidc.ScopesSupported, got.ScopesSupported)
	assert.Equal(t, oidc.ResponseTypesSupported, got.ResponseTypeSupported)
	assert.Equal(t, oidc.ResponseModesSupported, got.ResponseModesSupported)
	assert.Equal(t, oidc.GrantTypesSupported, got.GrantTypesSupported)
	assert.Equal(t, oidc.SubjectTypesSupported, got.SubjectTypesSupported)
	assert.Equal(t, oidc.IDTokenSigningAlgValuesSupported, got.IDTokenSigningAlgValuesSupported)
	assert.Equal(t, oidc.ClaimsSupported, got.ClaimsSupported)
	assert.Equal(t, privateJWKS, got.PrivateJWKS)
	assert.Equal(t, publicJWKS, got.PublicJWKS)
	assert.IsType(t, jwtSigner, got.JWTSigner) // Unfortunately, they are not equal.
}

func NewTestOP() *OP {
	jwks := NewJSONWebKeySetFromPEMs([]string{privateKeyPEM})
	return NewSimpleOP("http://example.com/op", jwks)
}
