package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizeResponse_ToQuery(t *testing.T) {
	testCases := []struct {
		authzRes AuthorizeResponse
		want     string
	}{
		{AuthorizeResponse{Code: "code1", State: ""}, "code=code1"},
		{AuthorizeResponse{Code: "code1", State: "xyz"}, "code=code1&state=xyz"},
	}
	for _, tc := range testCases {
		got := tc.authzRes.ToQuery()

		assert.Equal(t, tc.want, got)
	}
}
