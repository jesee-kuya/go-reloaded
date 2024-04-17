package main

import (
	"fmt"
	"os"

	goreloaded "goreloaded/texteditor"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments, try again")
		return
	} else if len(os.Args) > 3 {
		fmt.Println("More arguments than required, try again")
		return
	}
	sourceFile := os.Args[1]
	toFile := os.Args[2]

	file, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Could not read the file", err)
	}

	newContent := string(file)
	newContent = goreloaded.TextEditor(newContent)

	err = os.WriteFile(toFile, []byte(newContent), 0o644)
	if err != nil {
		fmt.Println("Error writing to file")
	}
}
