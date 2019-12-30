package collatzconjecture

import (
	"errors"
)

// CollatzConjecture function
func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("Invalid input")
	}

	return calculateConjecture(input), nil
}

func calculateConjecture(input int) int {
	if input == 1 {
		return 0
	}

	if input%2 == 0 {
		return 1 + calculateConjecture(input/2)
	}

	return 1 + calculateConjecture(input*3+1)
}
