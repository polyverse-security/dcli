package dcli


import (
	"os"

)

var (
	args     []string
)



var cliRunFunc = func(args []string, cc *CommandNode) {
	if len(args) < 2 {
		cc.Help()
		return
	}

	for _, sc := range cc.SubCommands() {
		if args[1] == sc.Name() {
			sc.Run(PassArgs(args, 2))
			return
		}
	}

	cc.Help()
}

var CLI = &CommandNode{
	N:       "bigbang-tools",
	D:       "toolkit for BigBang",
	U:       "bigbang-tools <subcommand> <args...>",
	RunFunc: cliRunFunc,
}

func Start() {
	CLI.Run(os.Args)
}
