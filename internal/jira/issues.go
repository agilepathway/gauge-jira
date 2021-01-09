package jira

import (
	"bytes"
	"fmt"

	"github.com/andygrunwald/go-jira"
)

type issues map[string]issue

func newIssues() issues {
	return make(map[string]issue)
}

func (i issues) addSpecs(specFilenames []string) {
	for _, filename := range specFilenames {
		i.addSpecToAllItsLinkedIssues(newSpec(filename))
	}
}

func (i issues) addSpecToAllItsLinkedIssues(spec spec) {
	for _, issueKey := range spec.issueKeys() {
		i.addSpecToIssue(spec, issueKey)
	}
}

func (i issues) addSpecToIssue(spec spec, issueKey string) {
	issue := i[issueKey]
	if issue.key == "" {
		issue.key = issueKey
	}

	issue.addSpec(spec)
	i[issueKey] = issue
}

func (i issues) publish() {
	var unpublishedIssues []issue

	jiraClient := jiraClient()

	for _, issue := range i {
		err := i.publishIssue(issue, jiraClient)
		if err != nil {
			unpublishedIssues = append(unpublishedIssues, issue)
			fmt.Printf("Failed to publish issue %s: %s", issue.key, err)
		}
	}

	switch len(i) - len(unpublishedIssues) {
	case 0:
		fmt.Println("No valid Jira specifications were found - so nothing to publish to Jira")
	case 1:
		fmt.Println("Published specifications to 1 Jira issue")
	default:
		fmt.Printf("Published specifications to %d Jira issues", len(i))
	}
}

func (i issues) publishIssue(issue issue, jiraClient *jira.Client) error {
	specs, err := issue.specsFormattedForJira()
	if err != nil {
		return err
	}

	req, err := jiraClient.NewRawRequest("PUT", fmt.Sprintf("rest/api/2/issue/%s", issue.key), bytes.NewBufferString(`{"update":{"description":[{"set": "`+specs+`"}]}}`)) //nolint:lll
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/json")

	_, err = jiraClient.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
