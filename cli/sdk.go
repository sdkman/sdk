package cli

import (
	"errors"
	"fmt"

	"github.com/sdkman/sdk/cmd"
	"github.com/sdkman/sdk/env"
	"github.com/sdkman/sdk/txt"
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
			version := env.GetLocalVersion(sdkmanDir)
			output = cmd.Version(version)
		case "selfupdate":
			output = cmd.Selfupdate("6.0.0", sdkmanDir)
		case "pull":
			fmt.Print(txt.Info("Pulling available version...  "))
			output, err = cmd.Pull(sdkmanApi)
			env.SetRemoteVersion("localhost", output, sdkmanDir)
			output = txt.InfoF("SDKMAN %s now available for installation...", output)
		default:
			return 1, errors.New(txt.ErrorF("No such command: %s", command))
		}

		fmt.Println(output)
		return 0, err
	}
}
