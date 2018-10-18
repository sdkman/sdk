package step

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"sdk/env"
)

var apiStub *httptest.Server

func theInstalledSdkmanVersionIs(version string) error {
	env.SetVersion(version, sdkmanDir)
	return nil
}

func theAvailableSdkmanVersionIs(version string) error {
	var response = fmt.Sprintf(`{"version": "%s"}`, version)
	apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
		w.WriteHeader(http.StatusOK)
	}))

	sdkmanApi = apiStub.URL

	return nil
}

func thePulledVersionStateIs(version string) error {

	file := filepath.Join(varDir, "localhost", "version")
	bytes, err := ioutil.ReadFile(file)
	check(err)

	actual := string(bytes)
	if actual != version {
		return fmt.Errorf("Expected %s but was %s", version, actual)
	}
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.Step(`^the installed sdkman version is "(.*)"$`, theInstalledSdkmanVersionIs)
	s.Step(`^the available sdkman version is "([^"]*)"$`, theAvailableSdkmanVersionIs)
	s.Step(`^the pulled version state is "([^"]*)"$`, thePulledVersionStateIs)
	s.AfterScenario(func(interface{}, error) {
		if apiStub != nil {
			fmt.Print("About to shut down apiStub server...\n")
			apiStub.Close()
			fmt.Print("Successfully shut down apiStub...\n")
		}
	})
}
