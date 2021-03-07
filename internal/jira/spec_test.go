package jira

import (
	"fmt"
	"os"
	"testing"
)

const (
	absolutePath       = "/home/vscode/workspace/gauge-jira/functional_tests/specs/"
	exampleFilename    = "example.spec"
	defaultExampleSpec = `
----
h1. This is the spec heading

h2. This is a scenario
`
	expectedDefaultExampleSpec = `
----
h1. This is the spec heading
[View or edit this spec in Git|%s/%s]


h2. This is a scenario
`
	linebreakBetweenHeadingAndScenario         = defaultExampleSpec
	expectedLinebreakBetweenHeadingAndScenario = expectedDefaultExampleSpec
	noLineBreakBetweenHeadingAndScenario       = `
----
h1. This is the spec heading
h2. This is a scenario
`
	expectedNoLineBreakBetweenHeadingAndScenario = `
----
h1. This is the spec heading
[View or edit this spec in Git|%s/%s]

h2. This is a scenario
`
	h1TextInScenario = `
----
h1. This is the spec heading
h2. This is a scenario which has the text h1. in it
`
	expectedh1TextInScenario = `
----
h1. This is the spec heading
[View or edit this spec in Git|%s/%s]

h2. This is a scenario which has the text h1. in it
`
)

var specTests = []struct { //nolint:gochecknoglobals
	input    string
	expected string
	filename string
}{
	{linebreakBetweenHeadingAndScenario, expectedLinebreakBetweenHeadingAndScenario, exampleFilename},
	{noLineBreakBetweenHeadingAndScenario, expectedNoLineBreakBetweenHeadingAndScenario, exampleFilename},
	{h1TextInScenario, expectedh1TextInScenario, exampleFilename},
	{defaultExampleSpec, expectedDefaultExampleSpec, "filename_with_the_word_specs.spec"},
}

//nolint:errcheck,gosec
func TestAddGitLinkAfterSpecHeading(t *testing.T) {
	for _, tt := range specTests {
		os.Setenv("GAUGE_SPEC_DIRS", absolutePath)
		os.Setenv("SPECS_GIT_URL", "https://github.com/agilepathway/gauge-jira/tree/master/functional-tests/specs")

		spec := Spec{absolutePath: absolutePath + tt.filename, markdown: ""}
		expected := fmt.Sprintf(tt.expected, os.Getenv("SPECS_GIT_URL"), tt.filename)
		actual := spec.addGitLinkAfterSpecHeading(tt.input)

		if expected != actual {
			t.Fatalf(`
	Expected
	%s
	
	but got:
	%s`, expected, actual)
		}
	}
}
