package goreloaded

import (
	"strings"
)

func Punctuate(str []string) string {
	punctuation := ""

	for i, v := range str {
		for _, val := range v {
			if val == '.' || val == ',' || val == '!' || val == '?' || val == ':' || val == ';' {
				punctuation += string(val)
			}
		}
		if len(punctuation) == 1 {
			str[i] = strings.Trim(str[i], punctuation)
			str[i] += punctuation
		} else if len(punctuation) > 1 {
			str[i-1] += punctuation
			str[i] = strings.Trim(str[i], punctuation)
		}
		punctuation = ""
	}
	for i, v := range str {
		if v == "." || v == "," || v == "!" || v == "?" || v == ":" || v == ";" {
			str[i-1] += v
			str[i] = ""
		}
	}

	finalStr := strings.Join(str, " ")
	return finalStr
}
