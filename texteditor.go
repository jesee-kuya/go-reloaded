package goreloaded

import (
	"regexp"
	"strings"
	"unicode"
)

// TextEditor returns a string that has been formated removing all format string
func TextEditor(newContent string) string {
	content := strings.Fields(newContent)
	for _, v := range newContent {
		if unicode.IsPunct(v) {
			newContent = Punctuate(content)
		}
	}
	test1 := regexp.MustCompile(`'\s* ([^']+)'`)
	newContent = test1.ReplaceAllString(newContent, " '$1'")
	newContent = Separate(newContent)
	content = strings.Fields(newContent)

	for _, v := range content {
		if v == "(cap)" || v == "(cap," {
			newContent = Cap(content)
			test := regexp.MustCompile(`\(cap,\s\d+\)`)
			newContent = test.ReplaceAllString(newContent, "")
			content = strings.Fields(newContent)
		} else if v == "(up)" || v == "(up," {
			newContent = Up(content)
			test := regexp.MustCompile(`\(up,\s\d+\)`)
			newContent = test.ReplaceAllString(newContent, "")
			content = strings.Fields(newContent)
		} else if v == "(low)" || v == "(low," {
			newContent = Low(content)
			test := regexp.MustCompile(`\(low,\s+\d+\)`)
			newContent = test.ReplaceAllString(newContent, "")
			content = strings.Fields(newContent)
		} else if v == "(hex)" {
			newContent = Hex(content)
			content = strings.Fields(newContent)
		} else if v == "(bin)" {
			newContent = Bin(content)
			content = strings.Fields(newContent)
		}
	}
	for _, v := range newContent {
		if unicode.IsPunct(v) {
			newContent = Punctuate2(content)
		}
	}
	content = strings.Fields(newContent)
	newContent = Vowel(content)
	return newContent
}
