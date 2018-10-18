package env

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// SetVersion writes the sdkman version
func SetVersion(version string, sdkmanDir string) {
	versionFile := filepath.Join(sdkmanDir, "var", "version")
	err := ioutil.WriteFile(versionFile, []byte(version), os.ModePerm)
	check(err)
}

// GetVersion reads the sdkman version
func GetVersion(sdkmanDir string) string {
	versionFile := filepath.Join(sdkmanDir, "var", "version")
	version, err := ioutil.ReadFile(versionFile)
	check(err)
	return string(version)
}

//SetRemoteVersion writes the sdkman remote version for a given domain
func SetRemoteVersion(domain string, version string, sdkmanDir string) {
	domainDir := filepath.Join(sdkmanDir, "var", domain)
	err := os.MkdirAll(domainDir, os.ModePerm)
	check(err)

	versionFile := filepath.Join(domainDir, "version")
	err = ioutil.WriteFile(versionFile, []byte(version), os.ModePerm)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
