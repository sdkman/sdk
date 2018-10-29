package env

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var sdkmanServiceEnv = "SDKMAN_API"

var sdkmanDirEnv = "SDKMAN_DIR"

var baseFolder = ".sdkman"

// SdkmanDir infers the sdkman directory based on the presence of the SDKMAN_DIR
// environment variable. If not present, the `$HOME/.sdkman` is used.
func SdkmanDir() string {
	sdkmanDir := os.Getenv(sdkmanDirEnv)
	if 0 < len(sdkmanDir) {
		return sdkmanDir
	}
	d, err := homedir.Dir()
	check(err)

	return filepath.Join(d, baseFolder)
}

// SdkmanApi infers the sdkman service host based on the presence of the SDKMAN_API
// environment variable. If not present, the https://api.sdkman.io/3 service is used.
func SdkmanApi() string {
	sdkmanApi := os.Getenv(sdkmanServiceEnv)
	if 0 < len(sdkmanApi) {
		return sdkmanApi
	}

	return "https://api.sdkman.io/3"
}
