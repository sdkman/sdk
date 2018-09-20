package cmd

import (
	"sdk/txt"
)

func Version(version string) string {
	return txt.InfoF("SDKMAN %s", version)
}
