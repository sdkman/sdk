package steps

import (
	"github.com/DATA-DOG/godog"
)

func theSdkmanVersionIs(_ string) error {
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.Step(`^the sdkman version is "(.*)"$`, theSdkmanVersionIs)
}
