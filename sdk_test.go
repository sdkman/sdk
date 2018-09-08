package main

import (
	"flag"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/kami-zh/go-capturer"
	"os"
	"strings"
	"testing"
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

func FeatureContext(s *godog.Suite) {
	s.Step(`^the internet is reachable$`, theInternetIsReachable)
	s.Step(`^the sdkman version is "(.*)"$`, theSdkmanVersionIs)
	s.Step(`^I enter "(.*)"$`, iEnter)
	s.Step(`^I see "(.*)"$`, iSee)
}

// for `go test` integration
var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
