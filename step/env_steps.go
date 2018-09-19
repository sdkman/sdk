package step

import (
	"github.com/DATA-DOG/godog"
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

func EnvContext(s *godog.Suite) {
	s.Step(`^the internet is reachable$`, theInternetIsReachable)
	s.Step(`^an initialised environment$`, anInitialisedEnvironment)
}
