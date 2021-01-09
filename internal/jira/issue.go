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

func (i *issue) publishSpecs() string {
	return json.Fmt(i.currentDescriptionWithExistingSpecsRemoved() +
		i.specsHeader() + i.jiraFmtSpecs() + i.specsFooter())
}

func (i *issue) specsHeader() string {
	return "----\nh2.Specification Examples\n"
}

func (i *issue) specsFooter() string {
	return "------------------------------\nEnd of specification examples\n----\n"
}

func (i *issue) jiraFmtSpecs() string {
	var jiraFmtSpecs string
	for _, spec := range i.specs {
		jiraFmtSpecs += spec.jiraFmt()
	}

	return jiraFmtSpecs
}

func (i *issue) currentDescriptionWithExistingSpecsRemoved() string {
	regexString := fmt.Sprintf("(?s)%s(.*)%s", i.specsHeader(), i.specsFooter())
	r := regexp.MustCompile(regexString)
	removed := r.ReplaceAllString(i.currentDescription(), "\n")

	if strings.TrimSpace(removed) == "" {
		return ""
	}

	return removed
}

func (i *issue) currentDescription() string {
	jiraClient := jiraClient()
	issue, _, _ := jiraClient.Issue.Get(i.key, nil)

	description := issue.Fields.Description
	if description == "" {
		return ""
	}

	return description + "\n"
}
