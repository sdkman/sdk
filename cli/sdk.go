package cli

import (
	"errors"
	"fmt"
	"sdk/cmd"
	"sdk/env"
	"sdk/txt"
)

// Sdk is the entry point used for delegating to the individual commands
// as well as performing and propagating error handling.
func Sdk(args []string, sdkmanDir string, sdkmanApi string) (int, error) {

	if len(args) == 0 {
		return 1, errors.New(txt.Error("No command specified."))

	} else {
		var command = args[0]
		var output string
		var err error

		switch command {
		case "version":
			version := env.GetVersion(sdkmanDir)
			output = cmd.Version(version)
		case "pull":
			fmt.Print(txt.Info("Pulling available version...  "))
			err, output = cmd.Pull(sdkmanApi)
			env.SetRemoteVersion("localhost", output, sdkmanDir)
			output = txt.InfoF("SDKMAN %s now available for installation...", output)
		default:
			return 1, errors.New(txt.ErrorF("No such command: %s", command))
		}

		fmt.Println(output)
		return 0, err
	}
}
