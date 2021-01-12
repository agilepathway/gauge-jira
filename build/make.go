/*----------------------------------------------------------------
 *  Copyright (c) ThoughtWorks, Inc.
 *  Licensed under the Apache License, Version 2.0
 *  See LICENSE in the project root for license information.
 *----------------------------------------------------------------*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	CGO_ENABLED = "CGO_ENABLED" //nolint:stylecheck,golint
)

const (
	dotGauge          = ".gauge"
	plugins           = "plugins"
	distros           = "distros"
	GOARCH            = "GOARCH" //nolint:golint
	GOOS              = "GOOS"
	X86               = "386"
	X86_64            = "amd64" //nolint:stylecheck
	DARWIN            = "darwin"
	LINUX             = "linux"
	WINDOWS           = "windows"
	bin               = "bin"
	newDirPermissions = 0755
	gauge             = "gauge"
	jira              = "jira"
	pluginJsonFile    = "plugin.json" //nolint:golint,stylecheck
)

func main() {
	flag.Parse()
	//nolint:gocritic
	if *install {
		updatePluginInstallPrefix()
		installPlugin(*pluginInstallPrefix)
	} else if *distro {
		createPluginDistro(*allPlatforms)
	} else if *test {
		runTests()
	} else {
		compile()
	}
}

func compile() {
	if *allPlatforms {
		compileAcrossPlatforms()
	} else {
		compileGoPackage(jira)
	}
}

func createPluginDistro(forAllPlatforms bool) {
	if forAllPlatforms {
		for _, platformEnv := range platformEnvs {
			setEnv(platformEnv)
			*binDir = filepath.Join(bin, fmt.Sprintf("%s_%s", platformEnv[GOOS], platformEnv[GOARCH]))
			fmt.Printf("Creating distro for platform => OS:%s ARCH:%s \n", platformEnv[GOOS], platformEnv[GOARCH])
			createDistro()
		}
	} else {
		createDistro()
	}

	log.Printf("Distributables created in directory => %s \n", bin)
}

func createDistro() {
	packageName := fmt.Sprintf("%s-%s-%s.%s", "gauge-"+jira, getPluginVersion(), getGOOS(), getArch())

	mirrorFile(pluginJsonFile, filepath.Join(getBinDir(), pluginJsonFile)) //nolint:errcheck,gosec
	os.Mkdir(filepath.Join(bin, distros), 0755)                            //nolint:errcheck,gosec
	createZipFromUtil(getBinDir(), packageName)
}

func createZipFromUtil(dir, name string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	os.Chdir(dir) //nolint:errcheck,gosec

	output, err := executeCommand("zip", "-r", filepath.Join("..", distros, name+".zip"), ".")
	fmt.Println(output)

	if err != nil {
		panic(fmt.Sprintf("Failed to zip: %s", err))
	}

	os.Chdir(wd) //nolint:errcheck,gosec
}

func isExecMode(mode os.FileMode) bool {
	return (mode & 0111) != 0
}

func mirrorFile(src, dst string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if sfi.Mode()&os.ModeType != 0 {
		log.Fatalf("mirrorFile can't deal with non-regular file %s", src)
	}

	dfi, err := os.Stat(dst)
	if err == nil &&
		isExecMode(sfi.Mode()) == isExecMode(dfi.Mode()) &&
		(dfi.Mode()&os.ModeType == 0) &&
		dfi.Size() == sfi.Size() &&
		dfi.ModTime().Unix() == sfi.ModTime().Unix() {
		// Seems to not be modified.
		return nil
	}

	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, newDirPermissions); err != nil {
		return err
	}

	df, err := os.Create(dst)
	if err != nil {
		return err
	}

	sf, err := os.Open(src) //nolint:gosec
	if err != nil {
		return err
	}
	defer sf.Close() //nolint:errcheck,gosec

	n, err := io.Copy(df, sf)
	if err == nil && n != sfi.Size() {
		err = fmt.Errorf("copied wrong size for %s -> %s: copied %d; want %d", src, dst, n, sfi.Size())
	}

	cerr := df.Close()

	if err == nil {
		err = cerr
	}

	if err == nil {
		err = os.Chmod(dst, sfi.Mode())
	}

	if err == nil {
		err = os.Chtimes(dst, sfi.ModTime(), sfi.ModTime())
	}

	return err
}

func mirrorDir(src, dst string) error {
	log.Printf("Copying '%s' -> '%s'\n", src, dst)
	err := filepath.Walk(src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		suffix, err := filepath.Rel(src, path)
		if err != nil {
			return fmt.Errorf("failed to find Rel(%q, %q): %v", src, path, err)
		}
		return mirrorFile(path, filepath.Join(dst, suffix))
	})

	return err
}

func runProcess(command string, arg ...string) {
	cmd := exec.Command(command, arg...) //nolint:gosec
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("Execute %v\n", cmd.Args)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func executeCommand(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...) //nolint:gosec
	bytes, err := cmd.Output()

	return strings.TrimSpace(fmt.Sprintf("%s", bytes)), err
}

func compileGoPackage(packageName string) { //nolint:unparam
	runProcess("go", "build", "-o", getGaugeExecutablePath(jira))
}

func getGaugeExecutablePath(file string) string {
	return filepath.Join(getBinDir(), getExecutableName(file))
}

func getExecutableName(file string) string {
	if getGOOS() == "windows" {
		return file + ".exe"
	}

	return file
}

func getBinDir() string {
	if *binDir != "" {
		return *binDir
	}

	return filepath.Join(bin, fmt.Sprintf("%s_%s", getGOOS(), getGOARCH()))
}

func getPluginVersion() string {
	pluginProperties, err := getPluginProperties(pluginJsonFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to get properties file. %s", err))
	}

	return pluginProperties["version"].(string)
}

func setEnv(envVariables map[string]string) {
	for k, v := range envVariables {
		os.Setenv(k, v) //nolint:errcheck,gosec
	}
}

func runTests() {
	if *verbose {
		runProcess("go", "test", "./...", "-v")
	} else {
		runProcess("go", "test", "./...")
	}
}

var test = flag.Bool("test", false, "Run the test cases")                                                                                               //nolint:gochecknoglobals,lll
var install = flag.Bool("install", false, "Install to the specified prefix")                                                                            //nolint:gochecknoglobals,lll
var pluginInstallPrefix = flag.String("plugin-prefix", "", "Specifies the prefix where the plugin will be installed")                                   //nolint:gochecknoglobals,lll
var distro = flag.Bool("distro", false, "Creates distributables for the plugin")                                                                        //nolint:gochecknoglobals,lll
var allPlatforms = flag.Bool("all-platforms", false, "Compiles or creates distributables for all platforms windows, linux, darwin both x86 and x86_64") //nolint:gochecknoglobals,lll
var binDir = flag.String("bin-dir", "", "Specifies OS_PLATFORM specific binaries to install when cross compiling")                                      //nolint:gochecknoglobals,lll
var verbose = flag.Bool("verbose", false, "Print verbose details")                                                                                      //nolint:gochecknoglobals,lll

var (
	platformEnvs = []map[string]string{ //nolint:gochecknoglobals
		map[string]string{GOARCH: X86_64, GOOS: DARWIN, CGO_ENABLED: "0"}, //nolint:gofmt
		map[string]string{GOARCH: X86, GOOS: LINUX, CGO_ENABLED: "0"},
		map[string]string{GOARCH: X86_64, GOOS: LINUX, CGO_ENABLED: "0"},
		map[string]string{GOARCH: X86, GOOS: WINDOWS, CGO_ENABLED: "0"},
		map[string]string{GOARCH: X86_64, GOOS: WINDOWS, CGO_ENABLED: "0"},
	}
)

func getPluginProperties(jsonPropertiesFile string) (map[string]interface{}, error) {
	pluginPropertiesJson, err := ioutil.ReadFile(jsonPropertiesFile) //nolint:golint,gosec,stylecheck
	if err != nil {
		fmt.Printf("Could not read %s: %s\n", filepath.Base(jsonPropertiesFile), err)
		return nil, err
	}

	var pluginJson interface{}                                                       //nolint:golint,stylecheck
	if err = json.Unmarshal([]byte(pluginPropertiesJson), &pluginJson); err != nil { //nolint:unconvert
		fmt.Printf("Could not read %s: %s\n", filepath.Base(jsonPropertiesFile), err)
		return nil, err
	}

	return pluginJson.(map[string]interface{}), nil
}

func compileAcrossPlatforms() {
	for _, platformEnv := range platformEnvs {
		setEnv(platformEnv)
		fmt.Printf("Compiling for platform => OS:%s ARCH:%s \n", platformEnv[GOOS], platformEnv[GOARCH])
		compileGoPackage(jira)
	}
}

func installPlugin(installPrefix string) {
	pluginInstallPath := filepath.Join(installPrefix, jira, getPluginVersion())
	mirrorDir(getBinDir(), pluginInstallPath)                                    //nolint:errcheck,gosec
	mirrorFile(pluginJsonFile, filepath.Join(pluginInstallPath, pluginJsonFile)) //nolint:errcheck,gosec
}

func updatePluginInstallPrefix() {
	if *pluginInstallPrefix == "" {
		if runtime.GOOS == "windows" {
			*pluginInstallPrefix = os.Getenv("APPDATA")
			if *pluginInstallPrefix == "" {
				panic(fmt.Errorf("failed to find AppData directory"))
			}

			*pluginInstallPrefix = filepath.Join(*pluginInstallPrefix, gauge, plugins)
		} else {
			userHome := getUserHome()
			if userHome == "" {
				panic(fmt.Errorf("failed to find User Home directory"))
			}
			*pluginInstallPrefix = filepath.Join(userHome, dotGauge, plugins)
		}
	}
}

func getUserHome() string {
	return os.Getenv("HOME")
}

func getArch() string {
	arch := getGOARCH()
	if arch == X86 {
		return "x86"
	}

	return "x86_64"
}

func getGOARCH() string {
	goArch := os.Getenv(GOARCH)
	if goArch == "" {
		return runtime.GOARCH
	}

	return goArch
}

func getGOOS() string {
	os := os.Getenv(GOOS)
	if os == "" {
		return runtime.GOOS
	}

	return os
}
