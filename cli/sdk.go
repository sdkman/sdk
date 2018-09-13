package cli

import (
	"errors"
	"fmt"
	"sdk/cmds"
)

func Sdk(args []string) (int, error) {

	if len(args) == 0 {
		return 1, errors.New("No command specified")

	} else {

		var command = args[0]
		var output string

		switch command {
		case "version":
			output, _ = cmds.Version()
		default:
			output = "No such command: " + command
		}

		fmt.Println(output)
		return 0, nil
	}
}
