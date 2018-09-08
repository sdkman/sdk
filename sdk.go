package main

import (
	"fmt"
	"os"
	"sdkman/cmds"
)

func main() {
	sdk(os.Args)
}

func sdk(args []string) error {
	var out string
	switch args[0] {
	case "version":
		out, _ = cmds.Version()
	default:
		out = "No such command"
	}

	fmt.Println(out)

	return nil
}
