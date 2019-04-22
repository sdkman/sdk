package cmd

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var sdkmanDirEnv = "SDKMAN_DIR"

func TestSelfupdateSuccess(t *testing.T)  {
	version := "6.0.0"

	sdkmanDir, err := ioutil.TempDir("/tmp", "sdkman-")
	check(err)

	os.Setenv(sdkmanDirEnv, sdkmanDir)

	actual := Selfupdate(version, sdkmanDir)
	assert.Contains(t, actual, "SDKMAN 6.0.0 now available for installation...")
	assert.Contains(t, actual, "Would you like to upgrade now? (Y/n):")
}
