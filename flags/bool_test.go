package flags

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBool(t *testing.T) {
	x := NewBoolFlag("example", "example int flag", true)
	y := GetBoolFlag("example")
	assert.Equal(t, x, y)
}

//this test is meant to test that nil is returned from GetBoolFlag() and Value() and doesn't panic
func TestNilBool(t *testing.T) {
	y := GetBoolFlag("example")
	assert.Nil(t, y)
}

func TestBoolFlag_Parse(t *testing.T) {
	x := NewBoolFlag("example", "example int flag", true)
	flagsBuffer = append(flagsBuffer, flagBuffer{
		name: "example",
		value: "true",
	})
	err := x.Parse()
	assert.NoError(t, err)
	assert.Equal(t, true, *x.V)

	x = NewBoolFlag("example", "example int flag", true)
	flagsBuffer = []flagBuffer{
		{name: "example", value: "1"},
	}
	err = x.Parse()
	assert.NoError(t, err)
	assert.Equal(t, true, *x.V)

		x = NewBoolFlag("example", "example int flag", true)
	flagsBuffer = []flagBuffer{
		{name: "example", value: "on"},
	}
	err = x.Parse()
	assert.NoError(t, err)
	assert.Equal(t, true, *x.V)
}