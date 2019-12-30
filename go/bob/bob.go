package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)
	hasSomething := len(trimmedRemark) != 0
	hasLetters, _ := regexp.MatchString(`[A-z]`, trimmedRemark)
	isQuestion, _ := regexp.MatchString(`\?$`, trimmedRemark)
	isAllCapitals := strings.ToUpper(trimmedRemark) == trimmedRemark

	if !hasSomething {
		return "Fine. Be that way!"
	}

	if hasLetters && isQuestion && isAllCapitals {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if hasLetters && isAllCapitals {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
