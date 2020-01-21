package lsproduct

import "errors"

const testVersion = 3

// LargestSeriesProduct function
func LargestSeriesProduct(str string, span int) (int, error) {
	if span == 0 {
		return 1, nil
	}
	if span < 0 {
		return 0, errors.New("")
	}
	if len(str) < span {
		return -1, errors.New("")
	}

	// hold the last N in a circular buffer so we can do a single pass
	cbuff := make([]int, span)
	cidx := 0
	product := 1
	max := 0

	for _, c := range str {

		if c < '0' || c > '9' {
			return 0, errors.New("")
		}
		v := int(c - '0')

		if v == 0 {
			// restart if we hit a zero
			cidx = 0
			product = 1
			continue
		}

		product *= v

		//hold the number so we can divide it back out later
		cbuff[cidx%span] = v
		cidx++

		if cidx < span {
			continue // not enough to check the product yet
		}

		if product > max {
			// the product of cbuff == product here
			max = product
		}

		//remove the head of cbuff from the product before the next iteration
		product = product / cbuff[cidx%span]
	}

	return max, nil
}
