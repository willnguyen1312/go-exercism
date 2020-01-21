package summultiples

// SumMultiples function
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		for _, d := range divisors {
			if d == 0 {
				continue
			}

			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
