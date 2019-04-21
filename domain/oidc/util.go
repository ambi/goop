package oidc

// IsSupportedResponseMode は responseMode という名前の Response Mode をサポートしているかどうかを返す。
func IsSupportedResponseMode(responseMode string) bool {
	for _, supportedResponseMode := range ResponseModesSupported {
		if responseMode == supportedResponseMode {
			return true
		}
	}
	return false
}

// IsSupportedResponseType は responseType という名前の Response Type をサポートしているかどうかを返す。
func IsSupportedResponseType(responseType string) bool {
	for _, supportedResponseType := range ResponseTypesSupported {
		if responseType == supportedResponseType {
			return true
		}
	}
	return false
}

// IsSupportedScope は scope という名前のスコープをサポートしているかどうかを返す。
func IsSupportedScope(scope string) bool {
	for _, supportedScope := range ScopesSupported {
		if scope == supportedScope {
			return true
		}
	}
	return false
}
