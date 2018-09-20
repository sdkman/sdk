package main

import (
	"fmt"
	"os"
	"sdk/cli"
	"sdk/env"
)

func main() {

	sdkmanDir := env.SdkmanDir()

	exit, err := cli.Sdk(os.Args[1:], sdkmanDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(exit)
}
