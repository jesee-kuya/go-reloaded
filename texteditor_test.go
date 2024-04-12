package goreloaded_test

import (
	"testing"

	"goreloaded"
)

var testcases = []struct {
	name     string
	expected string
	got      string
}{
	{"(hex)", "30 files were added", "1E (hex) files were added"},
	{"(bin)", "It has been 2 years", "It has been 10 (bin) years"},
}

func TestTextEditor(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			got := goreloaded.TextEditor(tc.got)

			if got != expected {
				t.Errorf("%v not %v", expected, got)
			}
		})
	}
}
