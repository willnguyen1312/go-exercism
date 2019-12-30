package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate function
func Abbreviate(s string) string {
	var result string
	var prev string

	for _, letter := range s {
		isLetter, _ := regexp.MatchString(`[A-Za-z]`, string(letter))
		if checkPrevLetter(prev) && isLetter {
			result += strings.ToUpper(string(letter))
		}
		prev = string(letter)
	}

	return result
}

func checkPrevLetter(letter string) bool {
	return letter == "" || letter == " " || letter == "-" || letter == "_"
}
