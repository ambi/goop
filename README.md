# goop

GOOP is an implementation of OpenID Provider written in Golang.

# References

- The OAuth 2.0 Authorization Framework: https://tools.ietf.org/html/rfc6749
  - OAuth 2.0 Form Post Response Mode: https://openid.net/specs/oauth-v2-form-post-response-mode-1_0.html
  - OAuth 2.0 Multiple Response Type Encoding Practices: https://openid.net/specs/oauth-v2-multiple-response-types-1_0.html
  - OAuth 2.0 Authorization Server Metadata: https://tools.ietf.org/html/rfc8414
  - OAuth 2.0 Security Best Current Practice (Draft): https://tools.ietf.org/html/draft-ietf-oauth-security-topics-12
- OpenID Connect 1.0: https://openid.net/specs/openid-connect-core-1_0.html
  - OpenID Connect Discovery 1.0: https://openid.net/specs/openid-connect-discovery-1_0.html
  - OpenID Connect Dynamic Client Registration 1.0: https://openid.net/specs/openid-connect-registration-1_0.html
  - OpenID Connect Session Management 1.0: https://openid.net/specs/openid-connect-session-1_0.html
  - OpenID Connect Front-Channel Logout 1.0: https://openid.net/specs/openid-connect-frontchannel-1_0.html
  - OpenID Connect Back-Channel Logout 1.0: https://openid.net/specs/openid-connect-backchannel-1_0.html
- JSON Web Token (JWT): https://tools.ietf.org/html/rfc7519
  - JSON Web Signature (JWS): https://tools.ietf.org/html/rfc7515
  - JSON Web Encryption (JWE): https://tools.ietf.org/html/rfc7516
  - JSON Web Key (JWK): https://tools.ietf.org/html/rfc7517
  - JSON Web Algorithms (JWA): https://tools.ietf.org/html/rfc7518

# curl test

```
% curl "http://localhost:8080/.well-known/openid-configuration"
% curl "http://localhost:8080/jwks"
% curl -v "http://localhost:8080/authorize?client_id=client-id1&redirect_uri=http://localhost/callback&response_type=code&scope=openid"
% curl -v -XPOST -d "client_id=client-id1" -d "client_secret=client-secret1" -d "code=$CODE" -d "redirect_uri=http://localhost/callback" -d "grant_type=authorization_code" "http://localhost:8080/token"
```

# Architecture

GOOP is designed in [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).
