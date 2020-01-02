package proverb

import (
	"fmt"
)

// Proverb function
func Proverb(rhyme []string) []string {
	result := []string{}
	length := len(rhyme)

	for index, item := range rhyme {
		if index == length-1 {
			result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", item, rhyme[index+1]))
		}
	}
	return result
}
