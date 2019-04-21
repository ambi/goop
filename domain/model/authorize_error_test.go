package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizeError_ToQuery(t *testing.T) {
	testCases := []struct {
		authzErr AuthorizeError
		want     string
	}{
		{AuthorizeError{StatusCode: 400, Message: "invalid request", State: ""}, "error=invalid+request"},
		{AuthorizeError{StatusCode: 400, Message: "invalid request", State: "xyz"}, "error=invalid+request&state=xyz"},
	}
	for _, tc := range testCases {
		got := tc.authzErr.ToQuery()

		assert.Equal(t, tc.want, got)
	}
}
