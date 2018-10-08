package step

import (
	"github.com/DATA-DOG/godog"
	"sdk/env"
)

func theInstalledSdkmanVersionIs(version string) error {
	env.SetVersion(version, sdkmanDir)
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.Step(`^the installed sdkman version is "(.*)"$`, theInstalledSdkmanVersionIs)
}
