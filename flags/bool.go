package flags

import "fmt"

func NewBoolFlag(name, description string, required bool) *BoolFlag {
	// create map if it doesn't exist
	if boolFlags == nil {
		boolFlags = make(map[string]*BoolFlag)
	}

	// create BoolFlag and assign to map
	var bFlag = &BoolFlag{N: name, D: description, V: nil, R: required}
	boolFlags[name] = bFlag

	return bFlag
}

// ====== BoolFlags ======
type BoolFlag struct {
	N string // name
	D string // description
	V *bool  // value pointer
	R bool   // required flag
}

var boolFlags map[string]*BoolFlag

func GetBoolFlag(name string) *BoolFlag {
	for k, v := range boolFlags {
		if k == name {
			return v
		}
	}
	return nil
}

// A V of nil means the V was not set.
func (f *BoolFlag) Value() *bool {
	if f == nil {
		return nil
	}
	return f.V
}

func (f *BoolFlag) Parse() error {
	for _, buf := range flagsBuffer {
		if buf.name == f.N {
			switch buf.value {
			case "true":
				v := true
				f.V = &v
				break
			case "false":
				v := false
				f.V = &v
				break
			default:
				return fmt.Errorf("flag %s expected a bool V but was given %s", f.N, buf.value)
			}
		}
	}
	return nil
}

func (f *BoolFlag) IsSet() bool {
	if f.V != nil {
		return true
	}
	return false
}

func (f *BoolFlag) Description() string {
	return f.D
}

func (f *BoolFlag) Name() string {
	return f.N
}

func (f *BoolFlag) Required() bool {
	return f.R
}
