package hamming

import (
	"errors"
)

// Distance function calculate the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("invalid input")
	}

	var result int
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			result++
		}
	}

	return result, nil
}
