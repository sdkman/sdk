package env

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSetVersion(t *testing.T) {
	sdkmanDir, err := ioutil.TempDir("/tmp", "sdkman-")
	check(err)
	defer os.RemoveAll(sdkmanDir)

	err = os.Mkdir(sdkmanDir+"/var", 0755)
	check(err)

	expected := "1.0.0"
	SetVersion(expected, sdkmanDir)

	versionFile := filepath.Join(sdkmanDir, "var", "version")
	content, err := ioutil.ReadFile(versionFile)
	check(err)

	actual := string(content)

	assert.Equal(t, expected, actual)
}

func TestGetVersion(t *testing.T) {

	sdkmanDir, err := ioutil.TempDir("/tmp", "sdkman-")
	check(err)
	defer os.RemoveAll(sdkmanDir)

	err = os.Mkdir(sdkmanDir+"/var", 0755)
	check(err)

	expected := "1.0.0"
	versionFile := filepath.Join(sdkmanDir, "var", "version")
	err = ioutil.WriteFile(versionFile, []byte(expected), os.ModePerm)
	check(err)

	actual := GetVersion(sdkmanDir)
	assert.Equal(t, expected, actual)
}

func TestSetRemoteVersion(t *testing.T) {
	domain := "localhost"
	expected := "1.0.0"

	sdkmanDir, err := ioutil.TempDir("/tmp", "sdkman-")
	check(err)
	defer os.RemoveAll(sdkmanDir)

	SetRemoteVersion(domain, expected, sdkmanDir)

	file := filepath.Join(sdkmanDir, "var", domain, "version")
	content, err := ioutil.ReadFile(file)
	check(err)

	actual := string(content)

	assert.Equal(t, expected, actual)
}
