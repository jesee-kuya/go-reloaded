package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// Up returns a string with uppercase letters after every occurences of (up) and (up, format index)
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

				if num < i {
					for a := num; a > 0; a-- {
						b := i - a
						newStr[b] = strings.ToUpper(newStr[b])
					}
				} else {
					fmt.Printf("The specified format index %d for up is larger than the length of the words before %d\n", num, i)
				}

			}
			n = ""
		}
	}
	finalStr := strings.Join(newStr, " ")
	return finalStr
}
