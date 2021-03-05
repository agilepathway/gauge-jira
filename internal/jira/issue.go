package jira

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/agilepathway/gauge-jira/internal/json"
)

const (
	specsHeaderMessage = "h2.Specification Examples"
	specsHeader        = "----\n----\n" + specsHeaderMessage + "\n"
	specsSubheader     = "h3.Do not edit these examples here.  Edit them using Gauge.\n"
	specsFooter        = "----\nEnd of specification examples\n----\n----\n"
)

type issue struct {
	specs []spec
	key   string
}

func (i *issue) addSpec(spec spec) {
	i.specs = append(i.specs, spec)
}

func (i *issue) specsFormattedForJira() (string, error) {
	currentDescriptionWithExistingSpecsRemoved, err := i.currentDescriptionWithExistingSpecsRemoved()
	if err != nil {
		return "", err
	}

	return json.Fmt(currentDescriptionWithExistingSpecsRemoved +
		specsHeader + specsSubheader + i.jiraFmtSpecs() + specsFooter), nil
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
	regexString := fmt.Sprintf("(?s)%s(.*)%s", specsHeader, specsFooter)
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
	return strings.Count(string(desc), specsHeaderMessage) < 2 //nolint:gomnd
}
