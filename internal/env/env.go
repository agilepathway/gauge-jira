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
		panic(fmt.Sprintf("%s environment variable not set", key))
	}

	return value
}
