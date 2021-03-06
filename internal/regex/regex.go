// Package regex implements simple regex utility functions
package regex

import (
	"regexp"
)

// CountMatches counts the number of matches of regex in s.
func CountMatches(s, regex string) int {
	matches := regexp.MustCompile(regex).FindAllStringIndex(s, -1)
	return len(matches)
}
