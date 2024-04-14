package goreloaded

import (
	"strconv"
	"strings"
)

// Up searches for all occurences
func Up(str []string) string {
	var newStr []string
	var n string
	for i, v := range str {
		if v == "(up)" {
			str[i-1] = strings.ToUpper(str[i-1])
		}
	}
	for _, v := range str {
		if v != "(up)" {
			newStr = append(newStr, v)
		}
	}
	for i, v := range newStr {
		if v == "(up," {
			for _, ch := range newStr[i+1] {
				if ch != ')' {
					n += string(ch)
				}
				num, _ := strconv.Atoi(n)

				for a := num; a > 0; a-- {
					b := i - a
					newStr[b] = strings.ToUpper(newStr[b])
				}
			}
			n = ""
		}
	}
	finalStr := strings.Join(newStr, " ")
	return finalStr
}
