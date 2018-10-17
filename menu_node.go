package dcli

import (
	"fmt"
	"strings"
)

// MenuNode is a node that provides a menu interface
type MenuNode struct {
	D string // Description
	N string // Name

	subCommands []DiscoveryNode
}

type MenuNodeInput struct {
	Name        string
	Description string
}

func NewMenuNode(input MenuNodeInput) *MenuNode {
	return &MenuNode{
		N: input.Name,
		D: input.Description,
	}
}

// Run will iterate through the subCommands nodes looking for the correct N and then pass the remaining arguments
// to the child node.
func (mn *MenuNode) Run(args []string) {
	UsageSlice = append(UsageSlice, mn.N)

	if len(args) < 1 {
		mn.Help()
		return
	}

	for _, child := range mn.subCommands {
		if args[0] == child.Name() {
			child.Run(args[1:])
			return
		}
	}

	mn.Help()
}

func (mn *MenuNode) AddSubCommand(child DiscoveryNode) {
	mn.subCommands = append(mn.subCommands, child)
}

func (mn *MenuNode) Name() string {
	return mn.N
}

func (mn *MenuNode) Description() string {
	return mn.D
}

func (mn *MenuNode) printDescription() {
	fmt.Println(Cyan(fmt.Sprintf("\nDescription:")))
	fmt.Printf("    %-15s\n", mn.D)
}

func (mn *MenuNode) printCommands() {
	if len(mn.subCommands) > 0 {
		fmt.Println(Cyan("\nAvailable commands:"))
		for _, sc := range mn.subCommands {
			fmt.Printf("    %-15s %15s\n", sc.Name(), Yellow(sc.Description()))
		}
	}
}

func (mn *MenuNode) Help() {
	fmt.Println(strings.Join(UsageSlice, " "), "<command> (<subcommand>) [flags]")
	mn.printDescription()
	mn.printCommands()
}
