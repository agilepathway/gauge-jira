package jira

import (
	"testing"
)

func TestSpecsHeaderContract(t *testing.T) {
	issue := issue{nil, ""}
	input := `
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

	expected := `
This is a description of an issue.

More description after the specs.`

	actual, _ := issue.removeSpecsFrom(input)

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
