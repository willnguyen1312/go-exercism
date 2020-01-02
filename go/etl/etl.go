package etl

import (
	"strings"
)

// Transform function
func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)

	for key, value := range input {
		for _, item := range value {
			result[strings.ToLower((item))] = key
		}
	}
	return result
}
