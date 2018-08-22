package dcli

import (
"fmt"

"github.com/fatih/color"
)

func PassArgs(args []string, consumed int) []string {
	if len(args)-consumed < 1 {
		return nil
	}
	return args[consumed:]
}

type Command interface {
	Run([]string)
	SubCommands() []Command
	Help()
	Name() string
	Description() string
	Usage() string
}

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	pink   = color.New(color.FgHiMagenta).SprintFunc()
)

type CommandNode struct {
	N           string // Name
	D           string // Description
	U           string // Usage
	subCommands []Command
	RunFunc     func([]string, *CommandNode)
	parent      string
}

func (cc *CommandNode) AddSubCommand(sub Command) {
	cc.subCommands = append(cc.subCommands, sub)
}

func (cc *CommandNode) SubCommands() []Command {
	return cc.subCommands
}

func (cc *CommandNode) Name() string {
	return cc.N
}

func (cc *CommandNode) Description() string {
	return cc.D
}

func (cc *CommandNode) Usage() string {
	return cc.U
}

func (cc *CommandNode) Help() {
	fmt.Printf("Usage: %s\n\n", cc.U)
	if len(cc.subCommands) > 0 {
		fmt.Println(cyan("Subcommands:"))
	}
	for _, sc := range cc.subCommands {
		fmt.Printf("    %-15s %15s\n", sc.Name(), yellow(sc.Description()))
	}
}

func (cc *CommandNode) Run(args []string) {
	cc.RunFunc(args, cc)
}

