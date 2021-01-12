package jira

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/agilepathway/gauge-jira/internal/json"
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
		i.specsHeader() + i.jiraFmtSpecs() + i.specsFooter()), nil
}

func (i *issue) specsHeader() string {
	return "----\n----\nh2.Specification Examples\nh3.Do not edit these examples here.  Edit them using Gauge.\n"
}

func (i *issue) specsFooter() string {
	return "------------------------------\nEnd of specification examples\n----\n----\n"
}

func (i *issue) jiraFmtSpecs() string {
	var jiraFmtSpecs string
	for _, spec := range i.specs {
		jiraFmtSpecs += spec.jiraFmt()
	}

	return jiraFmtSpecs
}

func (i *issue) currentDescriptionWithExistingSpecsRemoved() (string, error) {
	regexString := fmt.Sprintf("(?s)%s(.*)%s", i.specsHeader(), i.specsFooter())
	r := regexp.MustCompile(regexString)
	currentDescription, err := i.currentDescription()

	if err != nil {
		return "", err
	}

	removed := r.ReplaceAllString(currentDescription, "\n")

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

	description := issue.Fields.Description
	if description == "" {
		return "", nil
	}

	return description + "\n", nil
}
