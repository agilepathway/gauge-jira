// Package jira publishes Gauge specifications to Jira
package jira

import (
	"github.com/agilepathway/gauge-jira/internal/env"
	"github.com/agilepathway/gauge-jira/util"
	"github.com/andygrunwald/go-jira"
)

// PublishSpecs publishes the given Gauge specifications to Jira
func PublishSpecs(specs []Spec) {
	issues := newIssues()
	issues.addSpecs(specs)
	issues.publish()
}

func jiraClient() *jira.Client {
	jiraBaseURL := env.GetRequired("JIRA_BASE_URL")
	jiraUsername := env.GetRequired("JIRA_USERNAME")
	jiraToken := env.GetRequired("JIRA_TOKEN")

	transport := jira.BasicAuthTransport{
		Username: jiraUsername,
		Password: jiraToken,
	}

	jiraClient, err := jira.NewClient(transport.Client(), jiraBaseURL)
	util.Fatal("Error while creating Jira Client", err)

	return jiraClient
}
