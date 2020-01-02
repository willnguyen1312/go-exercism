package hamming

import (
	"errors"
)

// Distance function calculate the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	var result int

	if len(a) != len(b) {
		return result, errors.New("Invalid input")
	}

	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			result++
		}
	}

	return result, nil
}
