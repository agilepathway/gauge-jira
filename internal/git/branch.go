package git

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func discoverCurrentBranch() (string, error) {
	_, head, err := findHeadFile()
	if err != nil {
		return "", fmt.Errorf("there was a problem obtaining the HEAD file: %w", err)
	}

	currentBranch, err := discoverCurrentBranchFromHeadFile(head)
	if err != nil {
		return "", fmt.Errorf("there was a problem obtaining the current branch from the HEAD file: %w", err)
	}

	return currentBranch, nil
}

func findHeadFile() (string, string, error) {
	return findGitFile("HEAD")
}

func discoverCurrentBranchFromHeadFile(headPath string) (string, error) {
	if headPath == "" {
		return "", fmt.Errorf("no HEAD file defined")
	}

	headFile, err := ioutil.ReadFile(headPath) //nolint:gosec

	if err != nil {
		return "", fmt.Errorf("failed to load %s due to %s", headPath, err)
	}

	return abbreviatedHead(string(headFile))
}

func abbreviatedHead(fullHead string) (string, error) {
	if !strings.Contains(fullHead, "refs/heads/") {
		return "", fmt.Errorf("git is in detached HEAD state, HEAD is: %s", fullHead)
	}

	s := strings.TrimSpace(fullHead)
	s = strings.TrimPrefix(s, "ref:")
	s = strings.TrimSpace(s)

	return strings.TrimPrefix(s, "refs/heads/"), nil
}
