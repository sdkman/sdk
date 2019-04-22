package cmd

import (
	"github.com/sdkman/sdk/env"
	"github.com/sdkman/sdk/txt"
)

func Selfupdate(version string, sdkmanDir string) string {
	output := txt.InfoF("SDKMAN %s now available for installation...", version)
	output = output + txt.Info("Would you like to upgrade now? (Y/n):")
	env.SetLocalVersion(version, sdkmanDir)
	output = output + txt.Info("Successfully upgraded SDKMAN!")
	return txt.Info(output)
}
