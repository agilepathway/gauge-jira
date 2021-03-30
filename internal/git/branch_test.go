package git

import (
	"errors"
	"fmt"
	"testing"
)

var buildAbbreviatedHeadTests = []struct { //nolint:gochecknoglobals
	input    string
	expected string
}{
	{"ref: refs/heads/master", "master"},
	{"ref: refs/heads/main", "main"},
	{"ref: refs/heads/somefeaturebranch", "somefeaturebranch"},
	{"ref: refs/heads/using-as-separator", "using-as-separator"},
	{"ref: refs/heads/using_as_separator", "using_as_separator"},
	{"ref: refs/heads/using/slash/as/separator", "using/slash/as/separator"},
	{"ref: refs/heads/.startswithadot", ".startswithadot"},

	// just in case there are somehow newlines in the HEAD file
	{"\nref: refs/heads/master", "master"},
	{"\nref: refs/heads/master\n", "master"},
	{"ref: refs/heads/master\n", "master"},

	// just in case there are somehow extra spaces in the HEAD file
	{"ref:  refs/heads/master", "master"},
	{"ref:   refs/heads/master", "master"},
}

func TestAbbreviatedHead(t *testing.T) {
	for _, tt := range buildAbbreviatedHeadTests {
		actual, _ := abbreviatedHead(tt.input)
		if tt.expected != actual {
			t.Errorf("abbreviatedHead(%s): expected %s, actual %s", tt.input, tt.expected, actual)
		}
	}
}

func TestDetachedHead(t *testing.T) {
	detachedHead := "345c470bf286aa3ca8e8cb5d68361d4a685ba7d1"
	expectedErr := fmt.Errorf("git is in detached HEAD state, HEAD is: %s", detachedHead)
	res, err := abbreviatedHead(detachedHead)

	if (err == nil) || (res != "") {
		t.Errorf("Expected error for input %v but received %v", detachedHead, res)
	}

	if errors.Is(err, expectedErr) {
		t.Errorf("Unexpected error for input %v: %v", detachedHead, err)
	}
}
