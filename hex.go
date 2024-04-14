package goreloaded

import (
	"strconv"
	"strings"
)

// Hex searches for any occurence of (hex) in a slice of strings
// It then converts the previous value from a hexadecimal number to a decimal
// every word before the occurence of hex should be a hexadecimal number

func Hex(str []string) string {
	var newStr []string
	var new int64
	for i, v := range str {
		if v == "(hex)" {
			new, _ = strconv.ParseInt(str[i-1], 16, 64)
			s := strconv.FormatInt(new, 10)
			str[i-1] = s
		}
	}
	for _, v := range str {
		if v != "(hex)" {
			newStr = append(newStr, v)
		}
	}
	finalStr := strings.Join(newStr, " ")
	return finalStr
}
