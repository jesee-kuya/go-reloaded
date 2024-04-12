package main

import (
	"fmt"
	"os"
	"strings"

	"goreloaded"
)

func main() {
	sourceFile := os.Args[1]
	toFile := os.Args[2]

	file, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Could not read the file", err)
	}

	content := strings.Fields(string(file))
	newContent := goreloaded.Vowel(content)
	newContent = goreloaded.TextEditor(newContent)

	err = os.WriteFile(toFile, []byte(newContent), 0o644)
	if err != nil {
		fmt.Println("Error writing to file")
	}
}
