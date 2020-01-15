package diffsquares

// SquareOfSum function
func SquareOfSum(nth int) int {
	sum := 0
	for i := 1; i <= nth; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares function
func SumOfSquares(nth int) int {
	result := 0
	for i := 1; i <= nth; i++ {
		result += i * i
	}
	return result
}

// Difference function
func Difference(nth int) int {
	squareOfSum := SquareOfSum(nth)
	sumOfSquares := SumOfSquares(nth)

	if squareOfSum > sumOfSquares {
		return squareOfSum - sumOfSquares
	}
	return sumOfSquares - sumOfSquares
}
