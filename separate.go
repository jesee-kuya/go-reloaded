package goreloaded

import "regexp"

//Separate is used to search for occurence of the format strings eg (up) and (cap, 2)
// It then ensures there is a space before and after the format strings 

func Separate(str string) string {
	test1 := regexp.MustCompile(`\(\w+,\s+\d+\)`)
	str = test1.ReplaceAllString(str, " $0 ")
	test2 := regexp.MustCompile(`\(\w+\)`)
	str = test2.ReplaceAllString(str, " $0 ")
	return str
}
