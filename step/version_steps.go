package step

import (
	"github.com/DATA-DOG/godog"
	"os"
)

func theSdkmanVersionIs(version string) error {

	//open existing file in write-only append mode
	f, err := os.OpenFile(versionFile, os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	//clobber default version with `version`
	_, err = f.WriteAt([]byte(version), 0)
	check(err)
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.Step(`^the sdkman version is "(.*)"$`, theSdkmanVersionIs)
}
