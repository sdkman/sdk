package cli

import (
	"errors"
	"fmt"
	"sdk/cmds"
	"sdk/txt"
)

func Sdk(args []string) (int, error) {

	if len(args) == 0 {
		return 1, errors.New(txt.Error("No command specified."))

	} else {

		var command = args[0]
		var output string

		switch command {
		case "version":
			output, _ = cmds.Version()
		default:
			output = txt.Error("No such command: %s", command)
		}

		fmt.Println(output)
		return 0, nil
	}
}
