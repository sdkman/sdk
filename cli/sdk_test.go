package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdkWithNoArgs(t *testing.T) {
	noArgs := []string{}
	code, err := Sdk(noArgs, "", "")
	assert.Equal(t, 1, code, "Expected return code 1")
	assert.Error(t, err, "Expected error for no args invocation of Sdk()")
}

func TestSdkWithInvalidArgs(t *testing.T) {
	invalidArg := []string{"invalid"}
	code, err := Sdk(invalidArg, "", "")
	assert.Equal(t, 1, code, "Expected return code 1")
	assert.Error(t, err, "Expected error for invalid command invocation of Sdk()")
}
