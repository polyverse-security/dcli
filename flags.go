package dcli

var flags map[string]int

func RegisterFlag(name string, size int) {
	if flags == nil {
		flags = make(map[string]int)
	}
	flags[name] = size
}
