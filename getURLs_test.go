package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	cases := []struct {
		name        string
		inputURL    string
		inputBody   string
		expected    []string
		expectedErr error
	}{
		{
			name:     "relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="/path/two">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected:    []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
			expectedErr: nil,
		},
		{
			name:     "invalid html body",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html></>
			`,
			expected:    []string{},
			expectedErr: errors.New("error"),
		},
		{
			name:     "absolute URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="https://one.com/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://two.com/path/two">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected:    []string{"https://one.com/path/one", "https://two.com/path/two"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && tc.expectedErr == nil {
				t.Fatalf("Test %v - %s FAIL: expected error: %v, actual error: %v", i, tc.name, tc.expectedErr, err)
			} else if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Test %v - %s FAIL: expected: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
