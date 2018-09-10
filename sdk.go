package main

import (
	"fmt"
	"os"
	"sdk/cli"
)

func main() {
	err := cli.Sdk(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
