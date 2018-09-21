package env

import (
	"io/ioutil"
	"os"
)

var varFolder = "/var/version"

// SetVersion writes the sdkman version
func SetVersion(version string, sdkmanDir string) {

	//open existing file in write-only append mode
	versionFile := sdkmanDir + varFolder
	f, err := os.OpenFile(versionFile, os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	//clobber default version with `version`
	_, err = f.WriteAt([]byte(version), 0)
	check(err)
}

// GetVersion reads the sdkman version
func GetVersion(sdkmanDir string) string {
	version, err := ioutil.ReadFile(sdkmanDir + varFolder)
	check(err)
	return string(version)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
