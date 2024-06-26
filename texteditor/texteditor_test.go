package goreloaded_test

import (
	"testing"

	goreloaded "goreloaded/texteditor"
)

var testcases = []struct {
	name     string
	expected string
	got      string
}{
	{"(hex)", "30 files were added", "1e (hex) files were added"},
	{"(bin)", "It has been 2 years", "It has been 10 (bin) years"},
	{"(up)", "Ready, set, GO!", "Ready, set, go (up)!"},
	{"(low)", "I should stop shouting", "I should stop SHOUTING (low)"},
	{"(cap)", "Welcome to the Brooklyn Bridge", "Welcome to the Brooklyn bridge (cap)"},
	{"punctuation", "I was sitting over there, and then BAMM!!", "I was sitting over there ,and then BAMM !!"},
	{"punctuation combo", "I was thinking... You were right", "I was thinking ... You were right"},
	{"exclamation mark", "I am exactly how they describe me: 'awesome'", "I am exactly how they describe me: ' awesome '"},
	{"exclamation mark2", "As Elton John said: 'I am the most well-known homosexual in the world'", "As Elton John said: ' I am the most well-known homosexual in the world '"},
	{"vowel", "There it was. An amazing rock!", "There it was. A amazing rock!"},
	{"combo1", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.", "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."},
	{"HEX and BIN", "Simply add 66 and 2 and you will see the result is 68.", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."},
	{"vowel1", "There is no greater agony than bearing an untold story inside you.", "There is no greater agony than bearing a untold story inside you."},
	{"puctuation1", "Punctuation tests are... kinda boring, don't you think!?", "Punctuation tests are ... kinda boring ,don't you think !?"},
}

func TestTextEditor(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			got := goreloaded.TextEditor(tc.got)

			if got != expected {
				t.Errorf("%v NOT %v", expected, got)
			}
		})
	}
}
