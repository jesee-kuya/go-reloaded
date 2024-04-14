package goreloaded

import "strings"

// Vowel returns a strings with every instance of a before a word starting with a h or vowel turned to an
func Vowel(str []string) string {
	for i, v := range str {
		if v == "a" && i != len(str)-1 {
			for ind, val := range str[i+1] {
				if ind == 0 && (val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' || val == 'h') {
					str[i] = "an"
				}
			}
		} else if v == "A" && i != len(str)-1 {
			for ind, val := range str[i+1] {
				if ind == 0 && (val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' || val == 'h') {
					str[i] = "An"
				}
			}
		}
	}
	finalStr := strings.Join(str, " ")
	return finalStr
}
