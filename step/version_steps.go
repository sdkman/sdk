package step

import (
	"github.com/DATA-DOG/godog"
	"sdk/env"
)

func theSdkmanVersionIs(version string) error {
	env.SetVersion(version, sdkmanDir)
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.Step(`^the sdkman version is "(.*)"$`, theSdkmanVersionIs)
}
