package jira

import (
	"fmt"
	"testing"
)

const (
	exampleFilename    = "example.spec"
	specAbsolutePath   = "/home/vscode/workspace/gauge-jira/functional_tests/specs/" + exampleFilename
	specGitURL         = "https://github.com/agilepathway/gauge-jira/tree/master/functional-tests/specs/" + exampleFilename
	defaultExampleSpec = `
----
h1. This is the spec heading

h2. This is a scenario
`
	expectedDefaultExampleSpec = `
----
h1. This is the spec heading
[View or edit this spec in Git|%s]


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
[View or edit this spec in Git|%s]

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
[View or edit this spec in Git|%s]

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
		spec := Spec{
			path:   specAbsolutePath,
			gitURL: specGitURL}
		expected := fmt.Sprintf(tt.expected, specGitURL)
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
