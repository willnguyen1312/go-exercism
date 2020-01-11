package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram function
func IsIsogram(input string) bool {
	input = strings.ToLower(input)

	for index, character := range input {
		if unicode.IsLetter(character) && strings.ContainsRune(input[index+1:], character) {
			return false
		}
	}

	return true
}
