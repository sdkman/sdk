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

func TestSdkmanApiNoEnvVar(t *testing.T) {
	os.Unsetenv(sdkmanServiceEnv)

	expected := "https://api.sdkman.io/3"

	actual := SdkmanApi()

	assert.Equal(t, expected, actual, "no sdkman api default value when SDKMAN_API env var unset")
}

func TestSdkmanApiEnvVarSet(t *testing.T) {
	expected := "https://api.sdkman.io/2"
	os.Setenv(sdkmanServiceEnv, expected)

	actual := SdkmanApi()

	assert.Equal(t, expected, actual)
}
