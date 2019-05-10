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

func TestIntFlag_Parse(t *testing.T) {
	x := NewIntFlag("example", "example int flag", true)
	flagsBuffer = append(flagsBuffer, flagBuffer{
		name: "example",
		value: "1",
	})
	err := x.Parse()
	assert.NoError(t, err)
	assert.Equal(t, 1, *x.V)

	//not sure what the behavior of this should be. An int is a whole number, but this returns 1.02
	x = NewIntFlag("example", "example int flag", true)
	flagsBuffer = []flagBuffer{
		{ name: "example", value: "1.02", },
	}
	err = x.Parse()
	assert.NoError(t, err)
	assert.Equal(t, 1, *x.V)

	//and again here?
	x = NewIntFlag("example", "example int flag", true)
	flagsBuffer = []flagBuffer{
		{ name: "example", value: "0.1", },
	}
	err = x.Parse()
	assert.NoError(t, err)
	assert.Equal(t, 0, *x.V)
}