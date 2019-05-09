package flags

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewString(t *testing.T) {
	x := NewStringFlag("example", "example int flag", true)
	y := GetStringFlag("example")
	assert.Equal(t, x, y)
}

//this test is meant to test that nil is returned from GetStringFlag() and Value() and doesn't panic
func TestNilString(t *testing.T) {
	y := GetStringFlag("example").Value()
	assert.Nil(t, y)
}
