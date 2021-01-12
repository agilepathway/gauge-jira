package jira

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/agilepathway/gauge-jira/util"
)

type spec struct {
	filename string
	markdown string
}

func newSpec(filename string) spec {
	return spec{filename, readMarkdown(filename)}
}

func (s *spec) issueKeys() []string {
	return parseIssueKeys(s.markdown)
}

func (s *spec) jiraFmt() string {
	jiraFormatted := mdToJira(s.markdown)
	return "----\n" + s.downsizeHeadings(jiraFormatted)
}

func (s *spec) downsizeHeadings(input string) string {
	input = regexp.MustCompile(`h1\.`).ReplaceAllString(input, "h3.")
	input = regexp.MustCompile(`h2\.`).ReplaceAllString(input, "h4.")

	return input
}

func readMarkdown(filename string) string {
	specBytes, err := ioutil.ReadFile(filename) //nolint:gosec
	util.Fatal(fmt.Sprintf("Error while reading %s file", filename), err)

	return string(specBytes)
}
