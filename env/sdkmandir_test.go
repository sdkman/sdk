package env

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"os/user"
	"testing"
)

func TestSdkmanDirNoEnvVar(t *testing.T) {
	os.Unsetenv(sdkmanDirEnv)

	actual := SdkmanDir()

	u, err := user.Current()
	check(err)
	expected := u.HomeDir + "/" + baseFolder

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
