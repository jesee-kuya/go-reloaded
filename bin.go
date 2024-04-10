package goreloaded

import (
	"strconv"
	"strings"
)

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
