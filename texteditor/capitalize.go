package goreloaded

// Capitalize returns a string with the first letter as an uppercase letter
func Capitalize(str string) string {
	newStr := ""
	for i, v := range str {
		if i == 0 && (v >= 97 && v <= 122) {
			v -= 32
			newStr += string(v)
		} else {
			newStr += string(v)
		}
	}
	return newStr
}
