package cmd

import "github.com/sdkman/sdk/txt"

// Version formats the version output message
func Version(version string) string {
	return txt.InfoF("SDKMAN %s", version)
}
