package dcli

import (
	"os"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	pink   = color.New(color.FgHiMagenta).SprintFunc()
)

func Start(top *CommandNode) {
	top.Run(os.Args)
}

