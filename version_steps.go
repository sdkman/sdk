package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/kami-zh/go-capturer"
	"strings"
)

var actual string

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

	actual = strings.TrimSuffix(capturer.CaptureStdout(func() {
		sdk(args)
	}), "\n")

	return nil
}

func iSee(expected string) error {

	if actual != expected {
		return fmt.Errorf("expected %s but was %s", expected, actual)
	}
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.Step(`^the internet is reachable$`, theInternetIsReachable)
	s.Step(`^the sdkman version is "(.*)"$`, theSdkmanVersionIs)
	s.Step(`^I enter "(.*)"$`, iEnter)
	s.Step(`^I see "(.*)"$`, iSee)
}
