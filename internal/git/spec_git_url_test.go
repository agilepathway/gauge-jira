package git

import (
	"testing"
)

var buildGitWebURLTests = []struct { //nolint:gochecknoglobals
	input    string
	expected string
}{
	// HTTP/S
	{"http://github.com/example-user/example-repo", "http://github.com/example-user/example-repo"},
	{"http://github.com/example-user/example-repo.git", "http://github.com/example-user/example-repo"},
	{"http://github.com:8080/example-user/example-repo.git", "http://github.com:8080/example-user/example-repo"},
	{"http://example.com/example-user/example-repo", "http://example.com/example-user/example-repo"},
	{"http://git@example.com/example-user/example-repo", "http://example.com/example-user/example-repo"},
	{"http://user@example.com/example-user/example-repo", "http://example.com/example-user/example-repo"},
	{"https://github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},
	{"https://github.com/example-user/example-repo.git", "https://github.com/example-user/example-repo"},
	{"https://github.com:8080/example-user/example-repo.git", "https://github.com:8080/example-user/example-repo"},
	{"https://example.com/example-user/example-repo", "https://example.com/example-user/example-repo"},
	{"https://git@example.com/example-user/example-repo", "https://example.com/example-user/example-repo"},
	{"https://user@example.com/example-user/example-repo", "https://example.com/example-user/example-repo"},

	// scp-like
	{"git@github.com:example-user/example-repo.git", "https://github.com/example-user/example-repo"},
	{"git@github.com:example-user/example-repo", "https://github.com/example-user/example-repo"},
	{"git@github.com:8080:example-user/example-repo.git", "https://github.com:8080/example-user/example-repo"},
	{"git@example.com:example-user/example-repo.git", "https://example.com/example-user/example-repo"},
	{"user@github.com:example-user/example-repo.git", "https://github.com/example-user/example-repo"},

	// SSH transport protocol
	{"ssh://github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},
	{"ssh://github.com/example-user/example-repo.git", "https://github.com/example-user/example-repo"},
	{"ssh://github.com:8000/example-user/example-repo", "https://github.com:8000/example-user/example-repo"},
	{"ssh://git@github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},
	{"ssh://git@github.com:1234/example-user/example-repo", "https://github.com:1234/example-user/example-repo"},
	{"ssh://user@github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},
	{"git+ssh://github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},

	// Git transport protocol
	{"git://github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},
	{"git://github.com/example-user/example-repo.git", "https://github.com/example-user/example-repo"},
	{"git://github.com:8000/example-user/example-repo", "https://github.com:8000/example-user/example-repo"},
	{"git://git@github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},
	{"git://git@github.com:1234/example-user/example-repo", "https://github.com:1234/example-user/example-repo"},
	{"git://user@github.com/example-user/example-repo", "https://github.com/example-user/example-repo"},
}

func TestBuildGitWebURL(t *testing.T) {
	for _, tt := range buildGitWebURLTests {
		actual, _ := buildGitWebURL(tt.input)
		if tt.expected != actual {
			t.Errorf("buildGitWebURL(%s): expected %s, actual %s", tt.input, tt.expected, actual)
		}
	}
}
