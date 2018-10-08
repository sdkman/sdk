package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionOut(t *testing.T) {
	output := Version("1.0.0")
	assert.Equal(t, "SDKMAN 1.0.0", output)
}
