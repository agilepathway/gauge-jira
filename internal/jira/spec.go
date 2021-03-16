package jira

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/agilepathway/gauge-jira/internal/regex"
	"github.com/agilepathway/gauge-jira/util"
)

// Spec decorates a Gauge specification so it can be published to Jira.
type Spec struct {
	path     string // absolute path to the specification file, including the filename
	markdown string // the spec contents
	gitURL   string // the URL for the specs directory on e.g. GitHub, GitLab, Bitbucket
}

// NewSpec returns a new Spec object for the spec at the given absolute path
func NewSpec(absolutePath string, gitURL string) Spec {
	return Spec{absolutePath, readMarkdown(absolutePath), gitURL}
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
	if s.gitURL == "" {
		return spec
	}

	replacement := fmt.Sprintf("${1}\n%s\n", s.gitLinkInJiraFormat())

	return regex.ReplaceFirstMatch(spec, replacement, regexp.MustCompile(`(h1.*)`))
}

func (s *Spec) gitLinkInJiraFormat() string {
	return fmt.Sprintf("[View or edit this spec in Git|%s]", s.gitURL)
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
