package accumulate

// Accumulate function
func Accumulate(given []string, convert func(string) string) []string {
	result := make([]string, len(given))

	for index, item := range given {
		result[index] = convert(item)
	}

	return result
}
