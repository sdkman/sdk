package cli

import (
	"errors"
	"fmt"
	"sdk/cmd"
	"sdk/env"
	"sdk/txt"
)

func Sdk(args []string, sdkmanDir string) (int, error) {

	if len(args) == 0 {
		return 1, errors.New(txt.Error("No command specified."))

	} else {
		var command = args[0]
		var output string

		switch command {
		case "version":
			version, err := env.ReadVersion(sdkmanDir)
			check(err)
			output, _ = cmd.Version(string(version))
		default:
			return 1, errors.New(txt.ErrorF("No such command: %s", command))
		}

		fmt.Println(output)
		return 0, nil
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
