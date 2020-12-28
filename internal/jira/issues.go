package jira

import (
	"bytes"
	"fmt"
	"os"

	"github.com/agilepathway/gauge-jira/util"
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
	jiraClient := jiraClient()

	for _, issue := range i {
		i.publishIssue(issue, jiraClient)
	}
}

func (i issues) publishIssue(issue issue, jiraClient *jira.Client) {
	req, err := jiraClient.NewRawRequest("PUT", fmt.Sprintf("rest/api/2/issue/%s", issue.key), bytes.NewBufferString(`{"update":{"description":[{"set": "`+issue.jiraFmtSpecs()+`"}]}}`)) //nolint:lll
	util.Fatal("Error while creating Jira request %v", err)

	req.Header.Set("Content-type", "application/json")

	_, err = jiraClient.Do(req, nil)
	util.Fatal(fmt.Sprintf("Error while executing Jira request: %v", req), err)
}

func jiraClient() *jira.Client {
	base := os.Getenv("JIRA_BASE_URL")
	transport := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USERNAME"),
		Password: os.Getenv("JIRA_TOKEN"),
	}

	jiraClient, err := jira.NewClient(transport.Client(), base)
	util.Fatal("Error while creating Jira Client", err)

	return jiraClient
}
