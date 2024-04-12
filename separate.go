package goreloaded

import "regexp"

func Separate(str string) string {
	test1 := regexp.MustCompile(`\(\w+,\s+\d+\)`)
	str = test1.ReplaceAllString(str, " $0 ")
	test2 := regexp.MustCompile(`\(\w+\)`)
	str = test2.ReplaceAllString(str, " $0 ")
	return str
}
