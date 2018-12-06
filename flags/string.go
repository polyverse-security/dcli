package flags

import (
	"github.com/pkg/errors"
)

func NewStringFlag(name, description string, required bool) *StringFlag {
	// create map if it doesn't exist
	if stringFlags == nil {
		stringFlags = make(map[string]*StringFlag)
	}

	// create StringFlag and assign to map
	var sFlag = &StringFlag{N: name, D: description, V: nil, R: required}
	stringFlags[name] = sFlag

	return sFlag
}

// // ====== StringFlags ======
type StringFlag struct {
	N string  // name
	D string  // description
	V *string // value pointer, can use this to set default value
	R bool    // required flag
}

var stringFlags map[string]*StringFlag

func GetStringFlag(name string) *StringFlag {
	for k, v := range stringFlags {
		if k == name {
			return v
		}
	}
	return nil
}

func (f *StringFlag) Value() *string {
	if f == nil {
		return nil
	}
	return f.V
}

func (f *StringFlag) Parse() error {
	if f == nil {
		return errors.New("StringFlag is not set.")
	}
	for _, buf := range flagsBuffer {
		if buf.name == f.N {
			f.V = &buf.value
			stringFlags[f.N] = f
			break
		}
	}
	return nil
}

func (f *StringFlag) IsSet() bool {
	if f == nil {
		return false
	}
	if f.V != nil {
		return true
	}
	return false
}

func (f *StringFlag) Description() string {
	if f == nil {
		return ""
	}
	return f.D
}

func (f *StringFlag) Name() string {
	if f == nil {
		return ""
	}
	return f.N
}

func (f *StringFlag) Required() bool {
	return f.R
}
