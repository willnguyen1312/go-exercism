package anagram

import (
	"strings"
	"unicode"
)

type counts [26]int

func count(s string) (c counts) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			c[unicode.ToLower(r)-'a']++
		}
	}
	return
}

// Detect selects the correct result given a word and a list of possible anagrams.
func Detect(subject string, candidates []string) []string {
	result := []string{}
	sc := count(subject)
	for _, c := range candidates {
		if len(subject) != len(c) || strings.EqualFold(subject, c) {
			continue
		}
		cc := count(c)
		if sc == cc {
			result = append(result, c)
		}
	}
	return result
}
