package env

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {

	sdkmanDir, err := ioutil.TempDir("", "sdkman-")
	if err != nil {
		t.Fatalf("failed to create temporal directory (%#v)", err)
	}
	t.Logf("sdkmanDir = %s", sdkmanDir)
	defer (func(keep bool) {
		if !keep {
			_ = os.RemoveAll(sdkmanDirEnv)
			t.Logf("delete %s", sdkmanDir)
		}
	})(doKeepTemp())

	varPath := filepath.Join(sdkmanDir, varFile)
	// Create a directory to store the `varFile`.
	err = os.MkdirAll(filepath.Dir(varPath), 0755)
	if err != nil {
		t.Fatalf("failed to create directory \"%s\" (%#v)", filepath.Dir(varPath), err)
	}

	expected := "1.0.0"
	err = ioutil.WriteFile(varPath, []byte(expected), os.ModePerm)
	if err != nil {
		t.Fatalf("failed to create %s (%#v)", varPath, err)
	}
	actual := GetVersion(sdkmanDir)
	assert.Equal(t, expected, actual)
}
