// Package git provides Git utility functions.
// The implementation for these methods was copied
// from JX (Jenkins X CLI) and then modified slightly:
// https://github.com/jenkins-x/jx/blob/master/pkg/gits/git_cli.go
package git

import (
	"fmt"
	"io/ioutil"

	gitcfg "github.com/go-git/go-git/v5/config"
)

// findGitConfigDir tries to find the `.git` directory either in the current directory or in parent directories
func findGitConfigDir() (string, string, error) {
	return findGitFile("config")
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
