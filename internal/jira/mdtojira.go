package jira

import (
	"github.com/kalafut/m2j"
)

// mdToJira converts GitHub Flavored Markdown,
// which Gauge specifications are written in,
// into Jira's own format.
// https://github.github.com/gfm/
// https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa
func mdToJira(str string) string {
	return m2j.MDToJira(str)
}
