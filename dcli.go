package dcli

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	Pink   = color.New(color.FgHiMagenta).SprintFunc()
)

type DiscoveryNode interface {
	Run([]string)
	Help()
	Name() string
	Description() string
	Usage() string
}

func Start(top DiscoveryNode) {
	osArgs := os.Args[1:]
	flags := make(map[string]string)
	var nodeArgs []string

	for i := 0; i < len(osArgs); i++ {
		// look for flags
		if strings.HasPrefix(osArgs[i], "--") {
			flags[strings.TrimPrefix(osArgs[i], "--")] = osArgs[i+1]
			i++
			continue
		}
		nodeArgs = append(nodeArgs, osArgs[i])

	}

	top.Run(nodeArgs)
}
