package step

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"sdk/env"

	"github.com/DATA-DOG/godog"
)

const prefix = "sdkman-"

var sdkmanDir string
var varDir string

func theInternetIsReachable() error {
	//Internet availability not relevant yet..
	return nil
}

func anInitialisedEnvironment() error {

	var err error
	sdkmanDir, err = ioutil.TempDir("", prefix)
	if err != nil {
		return err
	}
	// fmt.Fprintln(os.Stderr, "sdkmanDir =", sdkmanDir)
	varDir = filepath.Join(sdkmanDir, "var")
	err = os.Mkdir(varDir, 0755)
	if err != nil {
		return err
	}
	env.SetVersion("0.0.1", sdkmanDir)

	return nil
}

func EnvContext(s *godog.Suite) {
	s.Step(`^the internet is reachable$`, theInternetIsReachable)
	s.Step(`^an initialised environment$`, anInitialisedEnvironment)
	s.AfterScenario(func(_ interface{}, _ error) {
		if 0 < len(sdkmanDir) {
			_ = os.RemoveAll(sdkmanDir)
			// fmt.Fprintln(os.Stderr, "removed ", sdkmanDir)
		}
		sdkmanDir = ""
	})
}
