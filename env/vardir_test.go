package env

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
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

	content, err := ioutil.ReadFile(sdkmanDir + varFile)
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
	err = ioutil.WriteFile(sdkmanDir+varFile, []byte(expected), os.ModePerm)
	check(err)

	actual := GetVersion(sdkmanDir)
	assert.Equal(t, expected, actual)
}
