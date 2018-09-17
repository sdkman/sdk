package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSdkWithNoArgs(t *testing.T) {
	code, err := Sdk([]string{})

	assert.Equal(t, 1, code, "Expected return code 1")
	assert.Error(t, err, "Expected error for no args invocation of Sdk()")
}

func TestSdkWithInvalidArgs(t *testing.T) {
	code, err := Sdk([]string{"invalid"})
	assert.Equal(t, 1, code, "Expected return code 1")
	assert.Error(t, err, "Expected error for invalid command invocation of Sdk()")
}
