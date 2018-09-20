package env

import (
	"io/ioutil"
	"os"
	"os/user"
)

var varFolder = "/var/version"

// Infers the sdkman directory based on the presence of the SDKMAN_DIR
// environment variable. If not present, we use `$HOME/.sdkman`.
func SdkmanDir() string {
	var sdkmanDir string
	if sdkmanDir = os.Getenv("SDKMAN_DIR"); sdkmanDir == "" {
		u, err := user.Current()
		check(err)

		sdkmanDir = u.HomeDir + "/.sdkman"
	}

	return sdkmanDir
}

// Write the sdkman version
func WriteVersion(version string, sdkmanDir string) {

	//open existing file in write-only append mode
	versionFile := sdkmanDir + varFolder
	f, err := os.OpenFile(versionFile, os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	//clobber default version with `version`
	_, err = f.WriteAt([]byte(version), 0)
	check(err)
}

// Read the sdkman version
func ReadVersion(sdkmanDir string) string {
	version, err := ioutil.ReadFile(sdkmanDir + varFolder)
	check(err)
	return string(version)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
