package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestSdkmanDir(t *testing.T) {
	// Prepare for environment variable operations.
	saved, ok := os.LookupEnv(sdkmanDirEnv)
	if ok {
		t.Logf("saving %s (= %s)", sdkmanDirEnv, saved)
	}
	defer (func() {
		if ok {
			_ = os.Setenv(sdkmanDirEnv, saved)
			t.Logf("restore %s to \"%s\"", sdkmanDirEnv, saved)
		}
	})()

	t.Run(fmt.Sprintf("Obtain sdkman home dir without %s", sdkmanDirEnv), func(t *testing.T) {
		os.Unsetenv(sdkmanDirEnv)

		actual := SdkmanDir()

		homeDir, err := homedir.Dir()
		if err != nil {
			t.Fatalf("failed to obtain the home directory (%#v)", err)
		}
		expected := filepath.Join(homeDir, baseFolder)

		assert.Equal(t, expected, actual, "no sdkman dir default value when SDKMAN_DIR env var unset")
	})
	t.Run(fmt.Sprintf("Obtain sdkman home dir under %s existed", sdkmanDirEnv), func(t *testing.T) {
		sdkmanDir, err := ioutil.TempDir("", "sdkman-")
		if err != nil {
			t.Fatalf("failed to create temporary directory (%#v)", err)
		}
		t.Logf("sdkmanDir = %s", sdkmanDir)
		defer (func(keep bool) {
			if !keep {
				_ = os.RemoveAll(sdkmanDir)
				t.Logf("%s deleted", sdkmanDir)
			}
		})(doKeepTemp())

		os.Setenv(sdkmanDirEnv, sdkmanDir)

		actual := SdkmanDir()

		expected := sdkmanDir

		assert.Equal(t, expected, actual, "sdkman dir value not read from SDKMAN_DIR env var")
	})
}

// doKeepTemp indicates tester wants to keep temporary outputs after testing.
func doKeepTemp() bool {
	f := os.Getenv("KEEP_TEMP")
	if len(f) == 0 {
		return false
	}
	f = strings.ToLower(f)
	switch f {
	case "y", "yes", "true", "on", "1":
		return true
	case "n", "no", "false", "off", "0":
		return false
	default:
		return false
	}
}
