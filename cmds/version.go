package cmds

import (
	"sdk/txt"
)

var version = "6.0.0"

func Version() (string, error) {
	out := txt.InfoF("SDKMAN %s", version)
	return out, nil
}
