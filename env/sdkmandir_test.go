package env

import (
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSdkmanDirNoEnvVar(t *testing.T) {
	os.Unsetenv(sdkmanDirEnv)

	actual := SdkmanDir()

	homeDir, err := homedir.Dir()
	check(err)
	expected := filepath.Join(homeDir, baseFolder)

	assert.Equal(t, expected, actual, "no sdkman dir default value when SDKMAN_DIR env var unset")
}

func TestSdkmanDirEnvVarSet(t *testing.T) {
	sdkmanDir, err := ioutil.TempDir("/tmp", "sdkman-")
	check(err)

	os.Setenv(sdkmanDirEnv, sdkmanDir)

	actual := SdkmanDir()

	expected := sdkmanDir

	assert.Equal(t, expected, actual, "sdkman dir value not read from SDKMAN_DIR env var")
}
