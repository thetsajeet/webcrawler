package main

import (
	"errors"
	"testing"
)

func TestGetArgs(t *testing.T) {
	cases := []struct {
		name        string
		input       []string
		expected    CrawlArgs
		expectedErr error
	}{
		{
			name:        "no args",
			input:       []string{},
			expected:    CrawlArgs{},
			expectedErr: errors.New(NO_ARG),
		},
		{
			name:        "too many args",
			input:       []string{"first.com", "2", "3", "4"},
			expected:    CrawlArgs{},
			expectedErr: errors.New(TOO_MANY_ARGS),
		},
		{
			name:  "1 arg",
			input: []string{"http://first.com"},
			expected: CrawlArgs{
				baseURL:        "http://first.com",
				maxConcurrency: DEF_MAX_CONCURRENCY,
				maxPages:       DEF_MAX_PAGES,
			},
			expectedErr: nil,
		},
		{
			name:  "2 arg",
			input: []string{"http://first.com", "3"},
			expected: CrawlArgs{
				baseURL:        "http://first.com",
				maxConcurrency: 3,
				maxPages:       DEF_MAX_PAGES,
			},
			expectedErr: nil,
		},
		{
			name:  "1 arg",
			input: []string{"http://first.com", "4", "3"},
			expected: CrawlArgs{
				baseURL:        "http://first.com",
				maxConcurrency: 4,
				maxPages:       3,
			},
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
