package env

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var sdkmanDirEnv = "SDKMAN_DIR"

var baseFolder = ".sdkman"

// SdkmanDir infers the sdkman directory based on the presence of the SDKMAN_DIR
// environment variable. If not present, we use `$HOME/.sdkman`.
func SdkmanDir() string {
	sdkmanDir := os.Getenv("SDKMAN_DIR")
	if 0 < len(sdkmanDir) {
		return sdkmanDir
	}
	d, err := homedir.Dir()
	check(err)

	return filepath.Join(d, baseFolder)
}
