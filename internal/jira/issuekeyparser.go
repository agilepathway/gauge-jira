package jira

import (
	"regexp"

	"github.com/agilepathway/gauge-jira/internal/unique"
)

/*
parseIssueKeys returns the set of Jira issue keys found in the given input string,
filtering out any duplicates.

Jira issue keys must be of the format:

         <project key>-<issue number>

where the project key must meet the following requirements:
- The first character must be a letter
- All letters used in the project key must be from the Modern Roman Alphabet and upper case
- Only letters, numbers or the underscore character can be used,
and where the issue number must be a number.

See: https://confluence.atlassian.com/adminjiraserver/changing-the-project-key-format-938847081.html
*/
func parseIssueKeys(input string) []string {
	r := regexp.MustCompile(`(([A-Z][A-Z_0-9]+)-[\d]+)`)
	return unique.Strings(r.FindAllString(input, -1))
}
