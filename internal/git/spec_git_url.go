package git

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// copied from https://github.com/go-git/go-git/blob/bf3471db54b0255ab5b159005069f37528a151b7/internal/url/url.go#L9
var scpLikeURIRegExp = regexp.MustCompile(`^(?:(?P<user>[^@]+)@)?(?P<host>[^:\s]+):(?:(?P<port>[0-9]{1,5})(?:\/|:))?(?P<path>[^\\].*\/[^\\].*)$`) //nolint:lll

// SpecGitURL gives the remote Git URL (e.g. on GitHub, GitLab, Bitbucket etc) for a spec file
// Returns "" if the remote Git URL could not be obtained
func SpecGitURL(absoluteSpecPath, projectRoot string) string {
	remoteGitURL, err := discoverRemoteGitURL()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	gitWebURL, err := buildGitWebURL(remoteGitURL)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	relativeSpecPath := strings.TrimPrefix(absoluteSpecPath, projectRoot)

	branch, err := discoverCurrentBranch()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return gitWebURL + "/blob/" + branch + toURLFormat(relativeSpecPath)
}

// buildGitWebURL constructs the publicly accessible Git web URL from a Git remote URL
func buildGitWebURL(remoteGitURI string) (string, error) {
	url, err := url.Parse(remoteGitURI)

	isStandardURL := err == nil && url != nil
	if isStandardURL {
		webURL := gitWebURLScheme(url.Scheme) + "://" + url.Host + strings.TrimSuffix(url.Path, ".git")
		return webURL, nil
	}

	if isSCPStyleURI(remoteGitURI) {
		_, host, port, path := findScpLikeComponents(remoteGitURI)
		webURL := "https://" + hostAndPort(host, port) + "/" + strings.TrimSuffix(path, ".git")

		return webURL, nil
	}

	return "", fmt.Errorf("could not parse Git URL %s", remoteGitURI)
}

func hostAndPort(host, port string) string {
	if port == "" {
		return host
	}

	return host + ":" + port
}

func isSCPStyleURI(input string) bool {
	return scpLikeURIRegExp.MatchString(input)
}

func findScpLikeComponents(uri string) (user, host, port, path string) {
	m := scpLikeURIRegExp.FindStringSubmatch(uri)
	return m[1], m[2], m[3], m[4]
}

func gitWebURLScheme(input string) string {
	if input == "http" {
		return input
	}

	return "https"
}

// toURLFormat converts any Windows path slashes to URL format (i.e. forward slashes)
func toURLFormat(input string) string {
	return strings.ReplaceAll(input, "\\", "/")
}
