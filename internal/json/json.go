// Package json provides utility JSON functions.
package json

import "fmt"

// Fmt formats a string in JSON format
func Fmt(input string) string {
	return removeOpeningAndClosingQuotes(fmt.Sprintf("%#v", input))
}

func removeOpeningAndClosingQuotes(input string) string {
	return input[1 : len(input)-1]
}
