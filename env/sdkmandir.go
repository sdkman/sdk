package env

import (
	"os"
	"os/user"
)

// SdkmanDir infers the sdkman directory based on the presence of the SDKMAN_DIR
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
