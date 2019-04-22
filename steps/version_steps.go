package steps

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/sdkman/sdk/env"
	"net/http"
	"net/http/httptest"
)

var apiStub *httptest.Server

func theInstalledSdkmanVersionIs(version string) error {
	env.SetLocalVersion(version, sdkmanDir)
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
	env.SetRemoteVersion("localhost", version, sdkmanDir)
	return nil
}

func theInstalledSdkmanVersionWasUpgradedTo(expected string) error {
	actual := env.GetLocalVersion(sdkmanDir)
	if actual != expected {
		return fmt.Errorf("expected %s but was %s", expected, actual)
	}
	return nil
}

func VersionFeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(i interface{}) {
		apiStub = nil
	})
	s.Step(`^the installed sdkman version is "(.*)"$`, theInstalledSdkmanVersionIs)
	s.Step(`^the available sdkman version is "([^"]*)"$`, theAvailableSdkmanVersionIs)
	s.Step(`^the pulled version state is "([^"]*)"$`, thePulledVersionStateIs)
	s.Step(`^the installed sdkman version was upgraded to "([^"]*)"$`, theInstalledSdkmanVersionWasUpgradedTo)
	s.AfterScenario(func(interface{}, error) {
		if apiStub != nil {
			fmt.Print("About to shut down apiStub server...\n")
			apiStub.Close()
			fmt.Print("Successfully shut down apiStub...\n")
		}
	})
}
