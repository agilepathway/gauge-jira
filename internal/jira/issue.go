package jira

import (
	"fmt"
)

type issue struct {
	specs []spec
	key   string
}

func (i *issue) addSpec(spec spec) {
	i.specs = append(i.specs, spec)
}

func (i *issue) jiraFmtSpecs() string {
	// nolint: godox
	// TODO: do not just format the first spec
	// TODO: this method is currently doing two things: jira formatting and json formatting
	//   - so it should probably be split into separate methods
	return i.removeOpeningAndClosingQuotes(fmt.Sprintf("%#v", i.specs[0].jiraFmt()))
}

func (i *issue) removeOpeningAndClosingQuotes(spec string) string {
	return spec[1 : len(spec)-1]
}
