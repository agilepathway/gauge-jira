// Package git provides Git utility functions.
// The implementation for these methods was copied
// from JX (Jenkins X CLI) and then modified slightly:
// https://github.com/jenkins-x/jx/blob/master/pkg/gits/git_cli.go
package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	gitcfg "github.com/go-git/go-git/v5/config"
)

// discoverRemoteGitURL discovers the remote git URL from the git configuration in the
// `.git` directory located in the current directory or parent directories
func discoverRemoteGitURL() (string, error) {
	_, gitConfig, err := findGitConfigDir()
	if err != nil {
		return "", fmt.Errorf("there was a problem obtaining the remote Git URL: %w", err)
	}

	remoteGitURL, err := discoverRemoteGitURLFromGitConfig(gitConfig)
	if err != nil {
		return "", fmt.Errorf("there was a problem obtaining the remote Git URL: %w", err)
	}

	return remoteGitURL, nil
}

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

// discoverRemoteGitURLFromGitConfig discovers the remote git URL from the given git configuration
func discoverRemoteGitURLFromGitConfig(gitConf string) (string, error) {
	cfg, err := parseGitConfig(gitConf)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal %s due to %s", gitConf, err)
	}

	remotes := cfg.Remotes

	if len(remotes) == 0 {
		return "", nil
	}

	rURL := getRemoteURL(cfg, "origin")

	if rURL == "" {
		rURL = getRemoteURL(cfg, "upstream")
	}

	return rURL, nil
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

// findGitConfigDir tries to find the `.git` directory either in the current directory or in parent directories
func findGitConfigDir() (string, string, error) {
	return findGitFile("config")
}

func findHeadFile() (string, string, error) {
	return findGitFile("HEAD")
}

func findGitFile(fileName string) (string, string, error) {
	var err error

	dir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	for {
		gitDir := filepath.Join(dir, ".git/"+fileName)
		exists, err := fileExists(gitDir)

		if err != nil {
			return "", "", err
		}

		if exists {
			return dir, gitDir, nil
		}

		dirPath := strings.TrimSuffix(dir, "/")
		if dirPath == "" {
			return "", "", nil
		}

		p, _ := filepath.Split(dirPath)

		if dir == "/" || p == dir {
			return "", "", nil
		}

		dir = p
	}
}

// fileExists checks if path exists and is a file
func fileExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return !fileInfo.IsDir(), nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, fmt.Errorf("failed to check if file exists %s %w", path, err)
}

func parseGitConfig(gitConf string) (*gitcfg.Config, error) {
	if gitConf == "" {
		return nil, fmt.Errorf("no GitConfDir defined")
	}

	cfg := gitcfg.NewConfig()
	data, err := ioutil.ReadFile(gitConf) //nolint:gosec

	if err != nil {
		return nil, fmt.Errorf("failed to load %s due to %s", gitConf, err)
	}

	err = cfg.Unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s due to %s", gitConf, err)
	}

	return cfg, nil
}

// getRemoteURL returns the remote URL from the given git config
func getRemoteURL(config *gitcfg.Config, name string) string {
	if config.Remotes != nil {
		return firstRemoteURL(config.Remotes[name])
	}

	return ""
}

func firstRemoteURL(remote *gitcfg.RemoteConfig) string {
	if remote != nil {
		urls := remote.URLs
		if len(urls) > 0 {
			return urls[0]
		}
	}

	return ""
}
