package flags

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInt(t *testing.T) {
	x := NewIntFlag("example", "example int flag", true)
	y := GetIntFlag("example")
	assert.Equal(t, x, y)
}

//this test is meant to test that nil is returned from GetIntFlag() and Value() and doesn't panic
func TestNilInt(t *testing.T) {
	y := GetIntFlag("example").Value()
	assert.Nil(t, y)
}
