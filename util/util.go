/*----------------------------------------------------------------
 *  Copyright (c) ThoughtWorks, Inc.
 *  Licensed under the Apache License, Version 2.0
 *  See LICENSE in the project root for license information.
 *----------------------------------------------------------------*/
//nolint:golint,stylecheck
package util

import (
	"fmt"
	"os"
)

var projectRoot string //nolint:gochecknoglobals

const (
	gaugeProjectRoot = "GAUGE_PROJECT_ROOT"
)

func init() { //nolint:gochecknoinits
	projectRoot = os.Getenv(gaugeProjectRoot)
	err := os.Chdir(projectRoot)

	if err != nil {
		fmt.Printf("Failed to Change working dir to project root %s: %s\n", projectRoot, err)
		os.Exit(1)
	}
}

func GetProjectRoot() string { //nolint:golint
	return projectRoot
}

func Fatal(message string, err error) { //nolint:golint
	if err != nil {
		fmt.Printf("%s. Error: %s", message, err.Error())
		os.Exit(1)
	}
}
