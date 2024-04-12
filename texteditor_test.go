package goreloaded_test

import (
	"testing"
)

var testcases = []struct {
	name     string
	got      string
	expected string
}{
	{"(hex)", "1E (hex) files were added", "30 files were added"},
	{"(bin)", "It has been 10 (bin) years", "It has been 2 years"},
}

func TestTextEditor(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			got := tc.got

			if got != expected {
				t.Errorf("%v not %v", expected, got)
			}
		})
	}
}
