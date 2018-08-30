package dcli

import "fmt"

// MenuNode is a node that provides a menu interface
type MenuNode struct {
	description string // Description
	name        string // Name
	runFunc     func([]string)
	usage       string // Usage

	subCommands []DiscoveryNode
}

type MenuNodeInput struct {
	Name        string
	Description string
	Usage       string
	RunFunc     func([]string)
}

func NewMenuNode(input MenuNodeInput) *MenuNode {
	return &MenuNode{
		name:        input.Name,
		description: input.Description,
		usage:       input.Usage,
		runFunc:     input.RunFunc,
	}
}

// Run will iterate through the subCommands nodes looking for the correct name and then pass the remaining arguments
// to the child node.
func (cc *MenuNode) Run(args []string) {
	if len(args) < 1 {
		cc.Help()
		return
	}

	for _, child := range cc.subCommands {
		if args[0] == child.Name() {
			child.Run(args[1:])
			return
		}
	}

	cc.Help()
}

func (cc *MenuNode) AddSubCommand(child DiscoveryNode) {
	cc.subCommands = append(cc.subCommands, child)
}

func (cc *MenuNode) Name() string {
	return cc.name
}

func (cc *MenuNode) Description() string {
	return cc.description
}

func (cc *MenuNode) Usage() string {
	return cc.usage
}

func (cc *MenuNode) Help() {
	fmt.Printf("Usage: %s\n\n", cc.usage)
	if len(cc.subCommands) > 0 {
		fmt.Println(cyan("Subcommands:"))
	}
	for _, sc := range cc.subCommands {
		fmt.Printf("    %-15s %15s\n", sc.Name(), yellow(sc.Description()))
	}
}
