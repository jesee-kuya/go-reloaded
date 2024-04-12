package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"goreloaded"
)

func main() {
	sourceFile := os.Args[1]
	// toFile := os.Args[2]

	file, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Could not read the file", err)
	}

	content := strings.Fields(string(file))
	var newContent string
	newContent = goreloaded.Vowel(content)
	content = strings.Fields(newContent)

	for _, v := range content {
		if v == "(cap)" || v == "(cap," {
			newContent = goreloaded.Cap(content)
			test := regexp.MustCompile(`\(cap,\s\d+\)`)
			newContent = test.ReplaceAllString(newContent, "")
			content = strings.Fields(newContent)
		} else if v == "(up)" || v == "(up," {
			newContent = goreloaded.Up(content)
			test := regexp.MustCompile(`\(up,\s\d+\)`)
			newContent = test.ReplaceAllString(newContent, "")
			content = strings.Fields(newContent)
		} else if v == "(low)" || v == "(low," {
			newContent = goreloaded.Low(content)
			test := regexp.MustCompile(`\(low,\s+\d+\)`)
			newContent = test.ReplaceAllString(newContent, "")
			content = strings.Fields(newContent)
		} else if v == "(hex)" {
			newContent = goreloaded.Hex(content)
			content = strings.Fields(newContent)
		} else if v == "(bin)" {
			newContent = goreloaded.Bin(content)
			content = strings.Fields(newContent)
		}
	}

	for _, v := range newContent {
		if unicode.IsPunct(v) {
			newContent = goreloaded.Punctuate(content)
		}
	}
	test1 := regexp.MustCompile(`'\s*([^']+)'`)
	newContent = test1.ReplaceAllString(newContent, " '$1'")

	fmt.Println(newContent)
}
