package env

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSetLocalVersion(t *testing.T) {
	sdkmanDir := mkdirs("var")
	defer os.RemoveAll(sdkmanDir)

	expected := "1.0.0"
	SetLocalVersion(expected, sdkmanDir)

	versionFile := filepath.Join(sdkmanDir, "var", "version")
	content, err := ioutil.ReadFile(versionFile)
	check(err)

	actual := string(content)

	assert.Equal(t, expected, actual)
}

func TestGetVersion(t *testing.T) {

	sdkmanDir := mkdirs("var")
	defer os.RemoveAll(sdkmanDir)

	expected := "1.0.0"
	versionFile := filepath.Join(sdkmanDir, "var", "version")
	err := ioutil.WriteFile(versionFile, []byte(expected), os.ModePerm)
	check(err)

	actual := GetLocalVersion(sdkmanDir)
	assert.Equal(t, expected, actual)
}

func TestSetRemoteVersion(t *testing.T) {
	domain := "localhost"
	expected := "1.0.0"

	sdkmanDir := mkdirs("var")
	defer os.RemoveAll(sdkmanDir)

	SetRemoteVersion(domain, expected, sdkmanDir)

	file := filepath.Join(sdkmanDir, "var", domain, "version")
	content, err := ioutil.ReadFile(file)
	check(err)

	actual := string(content)

	assert.Equal(t, expected, actual)
}

func TestGetRemoteVersion(t *testing.T) {
	domain := "localhost"

	sdkmanDir := mkdirs("var", domain)
	defer os.RemoveAll(sdkmanDir)

	expected := "1.0.0"
	versionFile := filepath.Join(sdkmanDir, "var", domain, "version")
	err := ioutil.WriteFile(versionFile, []byte(expected), os.ModePerm)
	check(err)

	actual := GetRemoteVersion(domain, sdkmanDir)
	assert.Equal(t, expected, actual)
}

func mkdirs(elems ...string) string {
	sdkmanDir, err := ioutil.TempDir("/tmp", "sdkman-")
	check(err)

	elemPath := strings.Join(elems, "/")
	directory := filepath.Join(sdkmanDir, elemPath)
	err = os.MkdirAll(directory, 0755)
	check(err)

	return sdkmanDir
}