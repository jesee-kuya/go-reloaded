package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// cap returns a string with capitalized words after every occurence of (cap) or (cap, number)

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

				if num < len(newStr) {
					if num > 0 {
						for a := num; a > 0; a-- {
							b := i - a
							newStr[b] = Capitalize(newStr[b])
						}
					} else if num < 0 {
						for a := num; a > 0; a++ {
							b := i - a
							newStr[b] = Capitalize(newStr[b])
						}
					}
				} else {
					fmt.Printf("The specified format index %d for cap is larger than the length of the words before %d\n", num, i)
				}

			}
			n = ""
		}
	}
	finalStr := strings.Join(newStr, " ")
	return finalStr
}
