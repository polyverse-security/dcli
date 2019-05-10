package flags

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArgs(t *testing.T) {
	args := []string{
		"--arg1", "value1",
		"--arg2", "value2",
	}
	values := []string{
		"arg1", "value1",
		"arg2", "value2",
	}
	_ = ParseFlags(args)
	for i, v := range flagsBuffer {
		assert.Equal(t, values[i], v)
	}
}
