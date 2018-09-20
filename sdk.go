package main

import (
	"fmt"
	"os"
	"sdk/cli"
)

func main() {
	var sdkmanDir string
	if sdkmanDir = os.Getenv("SDKMAN_DIR"); sdkmanDir == "" {
		sdkmanDir = "$HOME/.sdkman"
	}

	exit, err := cli.Sdk(os.Args[1:], sdkmanDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(exit)
}
