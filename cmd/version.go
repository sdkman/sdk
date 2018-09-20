package cmd

import (
	"sdk/txt"
)

func Version(version string) (string, error) {
	out := txt.InfoF("SDKMAN %s", version)
	return out, nil
}
