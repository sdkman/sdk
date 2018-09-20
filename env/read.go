package env

import (
	"io/ioutil"
	"os"
	"os/user"
)

// Read the sdkman version
func ReadVersion(sdkmanDir string) (string, error) {
	version, err := ioutil.ReadFile(sdkmanDir + "/var/version")
	return string(version), err
}

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
