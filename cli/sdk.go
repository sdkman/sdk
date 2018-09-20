package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sdk/cmd"
	"sdk/txt"
)

func Sdk(args []string, sdkmanDir string) (int, error) {

	version, err := ioutil.ReadFile(sdkmanDir + "/var/version")
	if err != nil {
		return 1, errors.New(txt.ErrorF("Could not read version file: %s", err))
	}

	if len(args) == 0 {
		return 1, errors.New(txt.Error("No command specified."))

	} else {
		var command = args[0]
		var output string

		switch command {
		case "version":
			output, _ = cmd.Version(string(version))
		default:
			return 1, errors.New(txt.ErrorF("No such command: %s", command))
		}

		fmt.Println(output)
		return 0, nil
	}
}
