package env

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var versionFilename = "version"

// SetLocalVersion writes the sdkman version
func SetLocalVersion(version string, sdkmanDir string) {
	writeVersion(version, sdkmanDir, "var")
}

//SetRemoteVersion writes the sdkman remote version for a given domain
func SetRemoteVersion(domain string, version string, sdkmanDir string) {
	writeVersion(version, sdkmanDir, "var", domain)
}

func writeVersion(version string, sdkmanDir string, pathTokens ...string) {
	relativePath := strings.Join(pathTokens, "/")
	domainDir := filepath.Join(sdkmanDir, relativePath)
	err := os.MkdirAll(domainDir, os.ModePerm)
	check(err)

	versionFile := filepath.Join(domainDir, versionFilename)
	err = ioutil.WriteFile(versionFile, []byte(version), os.ModePerm)
	check(err)
}

// GetLocalVersion reads the sdkman version
func GetLocalVersion(sdkmanDir string) string {
	return readVersion(sdkmanDir, "var")
}

//GetRemoteVersion reads the sdkman remote version for a given domain
func GetRemoteVersion(domain string, sdkmanDir string) string {
	return readVersion(sdkmanDir, "var", domain)
}

func readVersion(sdkmanDir string, pathTokens ...string) string {
	relativePath := strings.Join(pathTokens, "/")
	versionFile := filepath.Join(sdkmanDir, relativePath, versionFilename)
	version, err := ioutil.ReadFile(versionFile)
	check(err)
	return string(version)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
