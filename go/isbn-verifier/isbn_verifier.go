package isbn

// IsValidISBN function
func IsValidISBN(input string) bool {
	multiplier, sum := 10, 0

	for i, v := range input {
		if v >= '0' && v <= '9' {
			sum += int(v-'0') * multiplier
			multiplier--
		} else if i == len(input)-1 && v == 'X' {
			sum += 10
			multiplier--
		}
	}

	return sum%11 == 0 && multiplier == 0
}
