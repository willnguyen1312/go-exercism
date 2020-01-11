package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram function
func IsIsogram(input string) bool {
	normalizedInput := strings.ToLower(input)

	for index, character := range normalizedInput {
		if unicode.IsLetter(character) && strings.ContainsRune(normalizedInput[index+1:], character) {
			return false
		}
	}

	return true
}
