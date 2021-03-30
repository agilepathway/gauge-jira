package git

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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
