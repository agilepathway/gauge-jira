package jira

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/agilepathway/gauge-jira/internal/json"
	"github.com/agilepathway/gauge-jira/internal/regex"
)

const (
	specsHeaderMessage = "Specification Examples"
	// Jira sometimes adds or removes a space after the heading, so we need to cater for both scenarios
	specsHeaderMessageRegex = "h2.\\s*" + specsHeaderMessage
	specsHeader             = "\n----\n----\nh2." + specsHeaderMessage + "\n"
	specsHeaderRegex        = "\\s*----\\s*----\\s*" + specsHeaderMessageRegex + "\\s*"
	specsFooter             = "\n----\nEnd of specification examples\n----\n----\n"
	specsFooterRegex        = "\\s*----\\s*End of specification examples\\s*----\\s*----\\s*"
)

type issue struct {
	specs []Spec
	key   string
}

func (i *issue) addSpec(spec Spec) {
	i.specs = append(i.specs, spec)
}

func (i *issue) specsSubheader() string {
	return "h3.Edit these examples in Git (link is below), not here in Jira\n"
}

func (i *issue) specsFormattedForJira() (string, error) {
	currentDescriptionWithExistingSpecsRemoved, err := i.currentDescriptionWithExistingSpecsRemoved()
	if err != nil {
		return "", err
	}

	return json.Fmt(currentDescriptionWithExistingSpecsRemoved +
		specsHeader + i.specsSubheader() + i.jiraFmtSpecs() + specsFooter), nil
}

func (i *issue) jiraFmtSpecs() string {
	var jiraFmtSpecs string
	for _, spec := range i.specs {
		jiraFmtSpecs += spec.jiraFmt()
	}

	return jiraFmtSpecs
}

func (i *issue) currentDescriptionWithExistingSpecsRemoved() (string, error) {
	currentDescription, err := i.currentDescription()

	if err != nil {
		return "", err
	}

	return i.removeSpecsFrom(currentDescription)
}

func (i *issue) removeSpecsFrom(input string) (string, error) {
	regexString := fmt.Sprintf("(?s)%s(.*)%s", specsHeaderRegex, specsFooterRegex)
	r := regexp.MustCompile(regexString)

	removed := r.ReplaceAllString(input, "\n")

	if strings.TrimSpace(removed) == "" {
		return "", nil
	}

	return removed, nil
}

func (i *issue) currentDescription() (string, error) {
	jiraClient := jiraClient()
	issue, _, err := jiraClient.Issue.Get(i.key, nil)

	if err != nil {
		return "", err
	}

	desc := issue.Fields.Description
	if desc == "" {
		return "", nil
	}

	if description(desc).isValid() {
		return desc + "\n", nil
	}

	return "", fmt.Errorf("%[1]s is in an invalid state."+ //nolint:stylecheck
		"It contains more than one Gauge examples section, but there should only ever be one or none."+
		"Remove all Gauge example sections from %[1]s in Jira manually and then rerun the Gauge Jira plugin", i.key)
}

type description string

// isValid indicates if a Jira issue description is in a valid state for publishing Gauge specs to.
// An issue should only ever contain one Gauge examples section, currently (NB we may change this in
// the future if we cater for more than one source repo separately publishing their Gauge specs to the
// same Jira issue).
// A Jira issue could get into an invalid state either because of a manual edit, or because of a bug
// (hypothetically) in the Gauge Jira plugin itself which inadvertently led to a duplicate examples
// section instead of replacing the existing one.
func (desc description) isValid() bool {
	return regex.CountMatches(string(desc), specsHeaderRegex) < 2 //nolint:gomnd
}
