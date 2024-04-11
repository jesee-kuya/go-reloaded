package goreloaded

import "strings"

func Punctuate(str []string) string {
	var punctuation string
	for i, v := range str {
		for _, val := range v {
			if val == '.' || val == '!' || val == '?' || val == ',' || val == ':' {
				punctuation += string(val)
			}
			if string(v[0]) == punctuation {
				str[i-1] += punctuation
				str[i] = strings.Trim(str[i], punctuation)
				punctuation = ""
			}
		}
		punctuation = ""
	}
	finalStr := strings.Join(str, " ")
	return finalStr
}
