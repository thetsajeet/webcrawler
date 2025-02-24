package main

import (
	"errors"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	cases := []struct {
		name        string
		inputURL    string
		expected    string
		expectedErr error
	}{
		{
			name:        "remove scheme(protocol://)",
			inputURL:    "http://domain.com/path",
			expected:    "domain.com/path",
			expectedErr: nil,
		},
		{
			name:        "remove scheme and trailing /",
			inputURL:    "https://domain.com/path/",
			expected:    "domain.com/path",
			expectedErr: nil,
		},
		{
			name:        "reject relative paths",
			inputURL:    "domain.com/path",
			expected:    "",
			expectedErr: errors.New("will go wrong"),
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if err != nil && tc.expectedErr == nil {
				t.Fatalf("Test %d - %s FAIL: expected error: %v, actual error: %v", i, tc.name, tc.expectedErr, err)
			} else if actual != tc.expected {
				t.Fatalf("Test %d - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
