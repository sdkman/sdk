package steps

import (
	"github.com/DATA-DOG/godog"
	"github.com/sdkman/sdk/env"
	"io/ioutil"
	"os"
	"path/filepath"
)

var tmpFolder = "/tmp"
var prefix = "sdkman-"

var sdkmanApi string
var sdkmanDir string
var varDir string

func theInternetIsReachable() error {
	//Internet availability not relevant yet..
	return nil
}

func anInitialisedEnvironment() error {

	var err error
	sdkmanDir, err = ioutil.TempDir(tmpFolder, prefix)
	check(err)

	varDir = filepath.Join(sdkmanDir, "var")
	err = os.Mkdir(varDir, 0755)
	check(err)

	env.SetLocalVersion("0.0.1", sdkmanDir)

	return nil
}

func EnvContext(s *godog.Suite) {
	s.Step(`^the internet is reachable$`, theInternetIsReachable)
	s.Step(`^an initialised environment$`, anInitialisedEnvironment)

	//clean up old temp dir(s) before starting new test
	s.BeforeScenario(func(interface{}) {
		dirs, err := filepath.Glob(tmpFolder + "/" + prefix + "*")
		check(err)
		for _, d := range dirs {
			if err := os.RemoveAll(d); err != nil {
				panic(err)
			}
		}
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
