package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yandex/perforator/library/go/ptr"
)

func TestParseEndpoint(t *testing.T) {
	cases := []struct {
		name     string
		ep       string
		secure   *bool
		expected endpoint
	}{
		{
			name:   "AllSet",
			ep:     "localhost:8000",
			secure: ptr.T(false),
			expected: endpoint{
				host:   "localhost",
				port:   8000,
				secure: false,
			},
		},
		{
			name:   "SecureByDefault",
			ep:     "localhost:8000",
			secure: nil,
			expected: endpoint{
				host:   "localhost",
				port:   8000,
				secure: true,
			},
		},
		{
			name:   "DefaultSecurePort",
			ep:     "localhost",
			secure: ptr.T(true),
			expected: endpoint{
				host:   "localhost",
				port:   7618,
				secure: true,
			},
		},
		{
			name:   "DefaultInSecurePort",
			ep:     "localhost",
			secure: ptr.T(false),
			expected: endpoint{
				host:   "localhost",
				port:   7617,
				secure: false,
			},
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.name, func(t *testing.T) {
			endpoint, err := parseEndpoint(tcase.ep, tcase.secure)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tcase.expected, endpoint)
		})

	}
}
