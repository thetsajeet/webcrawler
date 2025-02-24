package main

import (
	"errors"
	"testing"
)

func TestGetArgs(t *testing.T) {
	cases := []struct {
		name        string
		input       []string
		expected    string
		expectedErr error
	}{
		{
			name:        "no args",
			input:       []string{},
			expected:    "",
			expectedErr: errors.New(NO_ARG),
		},
		{
			name:        "too many args",
			input:       []string{"first.com", "second.com"},
			expected:    "",
			expectedErr: errors.New(TOO_MANY_ARGS),
		}, {
			name:        "one arg",
			input:       []string{"http://first.com"},
			expected:    "http://first.com",
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetArgs(tc.input)
			if err != nil && tc.expectedErr == nil {
				t.Fatalf("Test %d - %s FAIL: expected error: %v, actual error: %v", i, tc.name, tc.expectedErr, err)
			} else if actual != tc.expected {
				t.Fatalf("Test %d - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
