package dcli

import "fmt"

// CommandNode is a node that executes a function.
type CommandNode struct {
	description      string // Description
	name             string // Name
	prompt           bool   // Require a prompt to execute the RunFunc
	runFunc          func([]string)
	usage            string // Usage
	requiredArgCount int    // the number of args this node is required to receive
}

func (cc *CommandNode) Name() string {
	return cc.name
}

func (cc *CommandNode) Description() string {
	return cc.description
}

func (cc *CommandNode) Usage() string {
	return cc.usage
}

func (cc *CommandNode) Help() {
	fmt.Printf("Usage: %s\n\n", cc.usage)
}

func (cc *CommandNode) Run(args []string) {
	cc.runFunc(args)
}
