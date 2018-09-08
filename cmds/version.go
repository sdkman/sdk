package cmds

import "fmt"

var version = "3.2.1"

func Version() (string, error) {
	out := fmt.Sprintf("SDKMAN %s", version)
	return out, nil
}
