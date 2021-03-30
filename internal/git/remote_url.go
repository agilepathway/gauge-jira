package git

import (
	"fmt"

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
