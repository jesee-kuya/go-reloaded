package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// low searches for any occurence of (low) or (low, number)
// Low returns a string with the previous value changed to lowercase.
func Low(str []string) string {
	var newStr []string
	var n string
	for i, v := range str {
		if v == "(low)" {
			str[i-1] = strings.ToLower(str[i-1])
		}
	}
	for _, v := range str {
		if v != "(low)" {
			newStr = append(newStr, v)
		}
	}
	for i, v := range newStr {
		if v == "(low," {
			for _, ch := range newStr[i+1] {
				if ch != ')' {
					n += string(ch)
				}
				num, _ := strconv.Atoi(n)

				if num < i {
					for a := num; a > 0; a-- {
						b := i - a
						newStr[b] = strings.ToLower(newStr[b])
					}
				} else {
					fmt.Printf("The specified format index %d for low is larger than the length of the words before %d\n", num, i)
					break
				}

			}
			n = ""
		}
	}
	finalStr := strings.Join(newStr, " ")
	return finalStr
}
