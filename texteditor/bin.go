package goreloaded

import (
	"strconv"
	"strings"
)

// The Bin function searches for any occurence of (bin) in a slice of strings
// It then converts the previous slice value from binary to decimal
// Every word before bin should be a binary number
// Bin returns a string

func Bin(str []string) string {
	var newStr []string
	for i, v := range str {
		if v == "(bin)" {
			new, _ := strconv.ParseInt(str[i-1], 2, 64)
			s := strconv.FormatInt(new, 10)
			str[i-1] = s
		}
	}
	for _, v := range str {
		if v != "(bin)" {
			newStr = append(newStr, v)
		}
	}
	finalStr := strings.Join(newStr, " ")
	return finalStr
}
