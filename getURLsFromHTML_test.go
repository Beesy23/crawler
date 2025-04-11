package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
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
		{
			name:      "no links",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><p>No links here</p></body></html>`,
			expected:  []string{},
		},
		{
			name:     "various URL formats",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html><body>
				<a href="https://absolute.com/path">Absolute</a>
				<a href="/path/with/slash">With slash</a>
				<a href="path/no/slash">No slash</a>
				<a href="?query=param">Query</a>
				<a href="#fragment">Fragment</a>
			</body></html>`,
			expected: []string{
				"https://absolute.com/path",
				"https://blog.boot.dev/path/with/slash",
				"https://blog.boot.dev/path/no/slash",
				"https://blog.boot.dev?query=param",
				"https://blog.boot.dev#fragment",
			},
		},
		{
			name:     "nested links",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html><body>
				<div>
					<p>
						<a href="/nested/deep">Deeply nested link</a>
					</p>
				</div>
				<a href="/sibling">Sibling link</a>
			</body></html>`,
			expected: []string{
				"https://blog.boot.dev/nested/deep",
				"https://blog.boot.dev/sibling",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
