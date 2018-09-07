package cli

import (
	"fmt"
)

func Sdk(args []string) error {
	var out string
	switch args[0] {
	case "version":
		out, _ = Version()
	default:
		out = "No such command"
	}

	fmt.Println(out)

	return nil
}
