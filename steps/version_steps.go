package steps

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/kami-zh/go-capturer"
	"sdk/cli"
	"strings"
)

var stdout string

var exitCode int

func theInternetIsReachable() error {
	//Internet availability not relevant yet..
	return nil
}

func anInitialisedEnvironment() error {
	//Initialise the ~/.sdkman folder here
	return nil
}

func theSdkmanVersionIs(_ string) error {
	return nil
}

func iEnter(command string) error {
	commandLine := strings.Split(command, " ")
	args := commandLine[1:]

	stdout = strings.TrimSuffix(capturer.CaptureStdout(func() {
		exitCode, _ = cli.Sdk(args)
	}), "\n")

	return nil
}

func iSee(expected string) error {

	if stdout != expected {
		return fmt.Errorf("expected %s but was %s", expected, stdout)
	}
	return nil
}

func theExitCodeIs(expected int) error {

	if exitCode != expected {
		return fmt.Errorf("expected %d but was %d", expected, exitCode)
	}
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.Step(`^the internet is reachable$`, theInternetIsReachable)
	s.Step(`^an initialised environment$`, anInitialisedEnvironment)
	s.Step(`^the sdkman version is "(.*)"$`, theSdkmanVersionIs)
	s.Step(`^I enter "(.*)"$`, iEnter)
	s.Step(`^I see "(.*)"$`, iSee)
	s.Step(`^the exit code is (\d+)$`, theExitCodeIs)
}
