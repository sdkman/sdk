package cmds

import "fmt"

var version = "6.0.0"

func Version() (string, error) {
	out := fmt.Sprintf("SDKMAN %s", version)
	return out, nil
}
