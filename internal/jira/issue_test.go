package jira

import (
	"testing"
)

const (
	inputWithNoSpaceAfterH2 = `
This is a description of an issue.
----
----
h2.Specification Examples
the spec
----
End of specification examples
----
----
More description after the specs.`
	inputWithSingleSpaceAfterH2 = `
This is a description of an issue.
----
----
h2. Specification Examples
the spec
----
End of specification examples
----
----
More description after the specs.`
	inputWithMultipleSpacesAfterH2 = `
This is a description of an issue.
----
----
h2.        Specification Examples
the spec
----
End of specification examples
----
----
More description after the specs.`
	inputWithLineBreaksBetweenRules = `
This is a description of an issue.
----

----
h2.        Specification Examples
the spec
----
End of specification examples
----

----
More description after the specs.`
	inputWithLineBreaksBeforeAndAfter = `
This is a description of an issue.


----
----
h2.        Specification Examples
the spec
----
End of specification examples
----
----


More description after the specs.`
	inputWithLineBreaksBetweenRulesAndSpecs = `
This is a description of an issue.
----
----

h2.        Specification Examples
the spec

----
End of specification examples
----
----
More description after the specs.`
	inputWithLineBreakAfterFooter = `
This is a description of an issue.
----
----
h2.        Specification Examples
the spec
----
End of specification examples

----
----
More description after the specs.`
)

var issueTests = []struct { //nolint:gochecknoglobals
	input string
}{
	{inputWithNoSpaceAfterH2},
	{inputWithSingleSpaceAfterH2},
	{inputWithMultipleSpacesAfterH2},
	{inputWithLineBreaksBetweenRules},
	{inputWithLineBreaksBeforeAndAfter},
	{inputWithLineBreaksBetweenRulesAndSpecs},
	{inputWithLineBreakAfterFooter},
}

func TestSpecsHeaderContract(t *testing.T) {
	for _, tt := range issueTests {
		issue := issue{nil, ""}
		expected := `
This is a description of an issue.
More description after the specs.`
		actual, _ := issue.removeSpecsFrom(tt.input)

		if expected != actual {
			t.Fatalf(`
	The contract for replacing specs in the issue description has changed. This would be a breaking change, as it would
	mean that existing users with Gauge specifications already published to Jira would not have these existing 
	specifications replaced when running the plugin.  Recommended solution therefore is to revert the change to the
	contract for replacing specs in the issue description.
	
	Expected
	%s
	
	but got:
	%s`, expected, actual)
		}
	}
}
