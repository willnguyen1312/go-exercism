package scrabble

import "unicode"

// Score function calculate the Scrabble score for the word
func Score(input string) int {
	var result int

	for _, letter := range input {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			result++
			break
		case 'D', 'G':
			result += 2
			break
		case 'B', 'C', 'M', 'P':
			result += 3
			break
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
			break
		case 'K':
			result += 5
			break
		case 'J', 'X':
			result += 8
			break
		case 'Q', 'Z':
			result += 10
			break
		}
	}

	return result
}
