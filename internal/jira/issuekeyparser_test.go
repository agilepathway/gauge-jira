package jira

import (
	"reflect"
	"testing"
)

var parseIssueKeysTests = []struct { //nolint:gochecknoglobals
	input    string
	expected []string
}{
	{"a string containing MYPROJECT-1", []string{"MYPROJECT-1"}},
	{"a string containing MYPROJECT-1, MYPROJECT-2", []string{"MYPROJECT-1", "MYPROJECT-2"}},
	{"a string containing MYPROJECT-1, MYPROJECT-1", []string{"MYPROJECT-1"}},
	{"a string containing MYPROJECT1-1, MYPROJECT2-1", []string{"MYPROJECT1-1", "MYPROJECT2-1"}},
	{"a string beginning with some-non-capital-lettersMYPROJECT-1", []string{"MYPROJECT-1"}},
	{"https://example.com/browse/MYPROJECT-1", []string{"MYPROJECT-1"}},
	{`[MYPROJECT-1](https://example.com/browse/MYPROJECT-1)`, []string{"MYPROJECT-1"}},
}

func TestParseIssueKeys(t *testing.T) {
	for _, tt := range parseIssueKeysTests {
		actual := parseIssueKeys(tt.input)
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("parseIssueKeys(%s): expected %s, actual %s", tt.input, tt.expected, actual)
		}
	}
}
