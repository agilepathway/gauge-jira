package jira

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/agilepathway/gauge-jira/internal/env"
	"github.com/agilepathway/gauge-jira/internal/regex"
	"github.com/agilepathway/gauge-jira/util"
)

// Spec decorates a Gauge specification so it can be published to Jira.
type Spec struct {
	absolutePath string
	markdown     string
}

// NewSpec returns a new Spec object for the spec at the given absolute path
func NewSpec(absolutePath string) Spec {
	return Spec{absolutePath, readMarkdown(absolutePath)}
}

func (s *Spec) issueKeys() []string {
	return parseIssueKeys(s.markdown)
}

func (s *Spec) jiraFmt() string {
	jiraFormattedSpec := mdToJira(s.markdown)
	jiraFormattedSpecWithGitLink := s.addGitLinkAfterSpecHeading(jiraFormattedSpec)

	return "----\n" + s.downsizeHeadings(jiraFormattedSpecWithGitLink)
}

func (s *Spec) addGitLinkAfterSpecHeading(spec string) string {
	replacement := fmt.Sprintf("${1}\n%s\n", s.gitLinkInJiraFormat())
	return regex.ReplaceFirstMatch(spec, replacement, regexp.MustCompile(`(h1.*)`))
}

func (s *Spec) gitLinkInJiraFormat() string {
	return fmt.Sprintf("[View or edit this spec in Git|%s]", s.gitURL())
}

func (s *Spec) gitURL() string {
	// ensure that we have the right number of slashes
	return strings.TrimSuffix(env.GetRequired("SPECS_GIT_URL"), "/") +
		"/" +
		strings.TrimPrefix(s.relativePath(), "/")
}

func (s *Spec) relativePath() string {
	return strings.TrimPrefix(s.absolutePath, os.Getenv("GAUGE_SPEC_DIRS"))
}

func (s *Spec) downsizeHeadings(spec string) string {
	spec = regexp.MustCompile(`h1\.`).ReplaceAllString(spec, "h3.")
	spec = regexp.MustCompile(`h2\.`).ReplaceAllString(spec, "h4.")

	return spec
}

func readMarkdown(absolutePath string) string {
	specBytes, err := ioutil.ReadFile(absolutePath) //nolint:gosec
	util.Fatal(fmt.Sprintf("Error while reading %s file", absolutePath), err)

	return string(specBytes)
}
