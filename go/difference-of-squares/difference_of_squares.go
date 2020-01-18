package diffsquares

// SquareOfSum function
func SquareOfSum(nth int) int {
	sum := (nth * (nth + 1)) / 2
	return sum * sum
}

// SumOfSquares function
func SumOfSquares(nth int) int {
	return (nth * (nth + 1) * (2*nth + 1)) / 6
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
