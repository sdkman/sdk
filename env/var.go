package env

import (
	"io/ioutil"
	"os"
)

var varFile = "/var/version"

// SetVersion writes the sdkman version
func SetVersion(version string, sdkmanDir string) {
	versionFile := sdkmanDir + varFile
	err := ioutil.WriteFile(versionFile, []byte(version), os.ModeAppend)
	check(err)
}

// GetVersion reads the sdkman version
func GetVersion(sdkmanDir string) string {
	version, err := ioutil.ReadFile(sdkmanDir + varFile)
	check(err)
	return string(version)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
