// Package env provides environment variables functionality.
package env

import (
	"fmt"
	"os"
)

// GetRequired returns an environment variable value or panics if not present.
func GetRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Aborting: %s is not set. Set it and try again. "+
			"See https://github.com/agilepathway/gauge-jira#plugin-setup", key))
	}

	return value
}
