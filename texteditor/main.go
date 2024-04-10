package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

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
	newContent1 := strings.Split(check, " ")
	check1 := goreloaded.Up(newContent1)
	newContent2 := strings.Split(check1, " ")
	check2 := goreloaded.Low(newContent2)
	newContent3 := strings.Split(check2, " ")
	check3 := goreloaded.Hex(newContent3)
	newContent4 := strings.Split(check3, " ")
	check4 := goreloaded.Bin(newContent4)
	newCheck4 := test.ReplaceAllString(check4, "")
	newContent5 := strings.Split(newCheck4, " ")
	check5 := goreloaded.Punctuate(newContent5)
	fmt.Println(check5)
}
