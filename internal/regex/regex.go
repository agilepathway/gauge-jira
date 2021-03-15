// Package regex implements simple regex utility functions
package regex

import (
	"regexp"
)

// CountMatches counts the number of matches of regex in s.
func CountMatches(s, pattern string) int {
	matches := regexp.MustCompile(pattern).FindAllStringIndex(s, -1)
	return len(matches)
}

// ReplaceFirstMatch returns a copy of src, replacing the first match of the
// Regexp pattern with the replacement string repl. Inside repl, $ signs are
// interpreted as in Expand, so for instance $1 represents the text of the
// first submatch.
func ReplaceFirstMatch(src, repl string, pattern *regexp.Regexp) string {
	isFirstMatch := true
	output := pattern.ReplaceAllStringFunc(src, func(match string) string {
		if isFirstMatch { // only do the replacing if it is the first match
			isFirstMatch = false
			return pattern.ReplaceAllString(match, repl)
		}
		return match // this is not the first match, so return the match unchanged
	})

	return output
}
