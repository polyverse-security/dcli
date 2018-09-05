package dcli

import "fmt"

// MenuNode is a node that provides a menu interface
type MenuNode struct {
	D string // Description
	N string // Name
	U string // Usage

	subCommands []DiscoveryNode
}

type MenuNodeInput struct {
	Name        string
	Description string
	Usage       string
}

func NewMenuNode(input MenuNodeInput) *MenuNode {
	return &MenuNode{
		N: input.Name,
		D: input.Description,
		U: input.Usage,
	}
}

// Run will iterate through the subCommands nodes looking for the correct N and then pass the remaining arguments
// to the child node.
func (mn *MenuNode) Run(args []string) {
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

func (mn *MenuNode) Usage() string {
	return mn.U
}

func (mn *MenuNode) Help() {
	fmt.Printf("Usage: %s\n\n", mn.U)
	if len(mn.subCommands) > 0 {
		fmt.Println(Cyan("Subcommands:"))
	}
	for _, sc := range mn.subCommands {
		fmt.Printf("    %-15s %15s\n", sc.Name(), Yellow(sc.Description()))
	}
}
