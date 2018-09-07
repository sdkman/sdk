package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/kami-zh/go-capturer"
	"log"
	"sdk/cli"
	"strings"
)

var Actual string

func theInternetIsReachable() error {
	//Internet availability not relevant yet..
	return nil
}

func theSdkmanVersionIs(version string) error {
	return nil
}

func iEnter(command string) error {
	commandLine := strings.Split(command, " ")
	args := commandLine[1:]

	out := capturer.CaptureStdout(func() {
		cli.Sdk(args)
	})

	Actual = strings.TrimSuffix(out, "\n")

	log.Printf("Actual: %s", Actual)

	return nil
}

func iSee(expected string) error {

	if expected != Actual {
		return fmt.Errorf("expected %s but was %s", expected, Actual)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the internet is reachable$`, theInternetIsReachable)
	s.Step(`^the sdkman version is "(.*)"$`, theSdkmanVersionIs)
	s.Step(`^I enter "(.*)"$`, iEnter)
	s.Step(`^I see "(.*)"$`, iSee)
}
