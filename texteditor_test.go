package goreloaded_test

import (
	"testing"

	"goreloaded"
)

func TestPunctuate(t *testing.T) {
	expected := "Punctuation tests are... kinda boring, don't you think!?"

	got := goreloaded.TextEditor("Punctuation tests are ... kinda boring ,don't you think ! ?")

	if got != expected {
		t.Errorf("Try again")
	}
}
