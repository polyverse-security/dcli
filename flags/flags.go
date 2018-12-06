// TODO when parsing flags, check against toggles to ensure there isn't a toggle and flag with the same name
package flags

import (
	"strings"
)

// Flags are parsed when the command is run, allowing you to register the same flag N more than once
type Flag interface {
	Name() string
	Description() string
	Required() bool
	IsSet() bool
	Parse() error
}

// flagBuffer is for storing flags at system start and holds them for parsing when a CommandNode is selected
type flagBuffer struct {
	name  string
	value string
}

var flagsBuffer []flagBuffer

// strip out all the flags and their values, the remaining args are for routing to CommandNodes
func ParseFlags(args []string) []string {
	var nodeArgs []string
	for i := 0; i < len(args); i++ {
		// look for flags
		if strings.HasPrefix(args[i], "--") {
			cleaned := strings.TrimPrefix(args[i], "--")
			// add it to the list of buffered flags. it'll be processed when the CommandNode runs
			if len(args) > i+1 {
				flagsBuffer = append(flagsBuffer, flagBuffer{name: cleaned, value: args[i+1]})
				i++
				continue
			}
		}
		nodeArgs = append(nodeArgs, args[i])
	}
	return nodeArgs
}

func GetFlag(name string) *string {
	for i := range flagsBuffer {
		if flagsBuffer[i].name == name {
			return &flagsBuffer[i].value
		}
	}
	return nil
}
