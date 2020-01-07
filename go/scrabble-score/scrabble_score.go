package scrabble

import "unicode"

// ScoreSwitch function calculate the Scrabble score for the word
func ScoreSwitch(input string) int {
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

var scrabbleScoreMap = map[rune]int{
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,
	'D': 2,
	'G': 2,
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,
	'K': 5,
	'J': 8,
	'X': 8,
	'Q': 10,
	'Z': 10,
}

// ScoreMap function calculate the Scrabble score for the word
func ScoreMap(input string) int {
	var result int

	for _, letter := range input {
		result += scrabbleScoreMap[unicode.ToUpper(letter)]
	}

	return result
}
