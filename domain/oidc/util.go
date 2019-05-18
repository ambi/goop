package oidc

// IsSupportedResponseMode returns whether a response type is supported or not.
func IsSupportedResponseMode(responseMode string) bool {
	for _, supportedResponseMode := range ResponseModesSupported {
		if responseMode == supportedResponseMode {
			return true
		}
	}
	return false
}

// IsSupportedResponseType returns whether a response type is supported or not.
func IsSupportedResponseType(responseType string) bool {
	for _, supportedResponseType := range ResponseTypesSupported {
		if responseType == supportedResponseType {
			return true
		}
	}
	return false
}

// IsSupportedScope returns whether a scope is supported or not.
func IsSupportedScope(scope string) bool {
	for _, supportedScope := range ScopesSupported {
		if scope == supportedScope {
			return true
		}
	}
	return false
}
