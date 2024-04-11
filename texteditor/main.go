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

	test := regexp.MustCompile(`\(\w+,\s\d+\)`)

	content := strings.Split(string(file), " ")
	fmt.Println(content)
	check := goreloaded.Cap(content)
	newContent1 := strings.Fields(check)
	check1 := goreloaded.Up(newContent1)
	newContent2 := strings.Fields(check1)
	check2 := goreloaded.Low(newContent2)
	newContent3 := strings.Fields(check2)
	check3 := goreloaded.Hex(newContent3)
	newContent4 := strings.Fields(check3)
	check4 := goreloaded.Bin(newContent4)
	newCheck4 := test.ReplaceAllString(check4, "")
	newContent5 := strings.Fields(newCheck4)
	for _, v := range newCheck4 {
		if unicode.IsPunct(v) {
			check4 = goreloaded.Punctuate(newContent5)
		}
	}

	test1 := regexp.MustCompile(`'\s*([^']+)'`)
	check5 := test1.ReplaceAllString(check4, " '$1'")

	fmt.Println(check5)
	newContent6 := strings.Fields(check5)
	check6 := goreloaded.Vowel(newContent6)
	fmt.Println(check6)
}
