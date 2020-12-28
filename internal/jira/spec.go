package jira

import (
	"fmt"
	"io/ioutil"

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
	// nolint:godox
	// TODO: implement this properly
	return []string{"JIRAGAUGE-1"}
}

func (s *spec) jiraFmt() string {
	return mdToJira(s.markdown)
}

func readMarkdown(filename string) string {
	specBytes, err := ioutil.ReadFile(filename) //nolint:gosec
	util.Fatal(fmt.Sprintf("Error while reading %s file", filename), err)

	return string(specBytes)
}
