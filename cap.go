package goreloaded

import (
	"strconv"
	"strings"
)

func Cap(str []string) string {
	var newStr []string
	var n string
	for i, v := range str {
		if v == "(cap)" && i != len(str) {
			str[i-1] = Capitalize(str[i-1])
		}
	}
	for _, v := range str {
		if v != "(cap)" {
			newStr = append(newStr, v)
		}
	}
	for i, v := range newStr {
		if v == "(cap," {
			for _, ch := range newStr[i+1] {
				if ch != ')' {
					n += string(ch)
				}
				num, _ := strconv.Atoi(n)

				for a := num; a > 0; a-- {
					b := i - a
					newStr[b] = Capitalize(newStr[b])
				}
			}
			n = ""
		}
	}
	finalStr := strings.Join(newStr, " ")
	return finalStr
}
