package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove the https scheme with trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove the http scheme with trailing slash",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "lowercase capital letters",
			inputURL: "https://blog.BOOT.DEV/pATH",
			expected: "blog.boot.dev/path",
		},
		{
			name:          "handler invalid URL",
			inputURL:      "://invalidURL",
			expected:      "",
			errorContains: "couldn't parse URL",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				assertErrorMessage(t, i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				assertErrorMessage(t, i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none", i, tc.name, tc.errorContains)
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: \nexpected URL: %v, \nactual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative paths",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s, FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				assertErrorMessage(t, i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: \nexpected: %v, \nactual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func assertErrorMessage(t testing.TB, i int, name string, err error) {
	t.Helper()
	t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, name, err)
}
