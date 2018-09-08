package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"sdk/cmds"
)

func main() {
	err := sdk(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sdk(args []string) error {

	if len(args) == 0 {
		return errors.New("No command specified")

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
		return nil
	}
}
