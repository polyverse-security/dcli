package flags

import (
	"fmt"
	"strconv"
)

func NewIntFlag(name, description string, required bool) *IntFlag {
	// create map if it doesn't exist
	if intFlags == nil {
		intFlags = make(map[string]*IntFlag)
	}

	// create BoolFlag and assign to map
	var iFlag = &IntFlag{N: name, D: description, V: nil, R: required}
	intFlags[name] = iFlag

	return iFlag
}

// ====== IntFlags ======
type IntFlag struct {
	N string // name
	D string // description
	V *int   // value pointer
	R bool   // required flag
}

var intFlags map[string]*IntFlag

func GetIntFlag(name string) *IntFlag {
	for k, v := range intFlags {
		if k == name {
			return v
		}
	}
	return nil
}

// A value of nil means the value was not set.
func (f *IntFlag) Value() *int {
	if f == nil {
		return nil
	}
	return f.V
}

func (f *IntFlag) Parse() error {
	for _, buf := range flagsBuffer {
		if buf.name == f.N {
			i, err := strconv.Atoi(buf.value)
			if err != nil {
				return fmt.Errorf("flag %s expected an int but was given %s", f.N, buf.value)
			}
			f.V = &i
			intFlags[f.N] = f
			break
		}
	}
	return nil
}

func (f *IntFlag) IsSet() bool {
	if f.V != nil {
		return true
	}
	return false
}

func (f *IntFlag) Description() string {
	return f.D
}

func (f *IntFlag) Name() string {
	return f.N
}

func (f *IntFlag) Required() bool {
	return f.R
}
