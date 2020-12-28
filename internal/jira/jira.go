// Package jira publishes Gauge specifications to Jira
package jira

// PublishSpecs publishes the given Gauge specifications to Jira
func PublishSpecs(specFilenames []string) {
	issues := newIssues()
	issues.addSpecs(specFilenames)
	issues.publish()
}
