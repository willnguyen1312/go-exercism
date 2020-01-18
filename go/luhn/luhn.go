package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid function
func Valid(input string) bool {
	normalizedInput := strings.Replace(input, " ", "", -1)

	if len(normalizedInput) <= 1 {
		return false
	}

	// Map string to a list of number
	numberList := make([]int, len(normalizedInput))
	for index, letter := range normalizedInput {
		if !unicode.IsDigit(letter) {
			return false
		}

		number, _ := strconv.Atoi(string(letter))
		numberList[index] = number
	}

	// Start doubling numberList from the right
	for index := len(numberList) - 2; index >= 0; index -= 2 {
		doubled := (numberList[index] * 2)
		if doubled > 9 {
			doubled -= 9
		}
		numberList[index] = doubled
	}

	sum := 0
	for _, number := range numberList {
		sum += number
	}

	return sum%10 == 0
}
