package dcli

import (
	"fmt"

	"github.com/ianchildress/dcli/flags"
)

// CommandNode is a node that executes a function.
type CommandNode struct {
	D       string // Description
	N       string // Name
	U       string // Description
	Prompt  bool   // Require a prompt to execute the RunFunc
	RunFunc func() // the function to run when this node is called
	flags   []flags.Flag
}

func (cn *CommandNode) Name() string {
	return cn.N
}

func (cn *CommandNode) Description() string {
	return cn.D
}

func (cn *CommandNode) Usage() string {
	return cn.U
}

func (cn *CommandNode) Help() {
	fmt.Println(Cyan(fmt.Sprintf("\nUsage:")))
	fmt.Printf("    %-15s\n", cn.U)
	fmt.Println(Cyan(fmt.Sprintf("\nFlags:")))
	for _, f := range cn.flags {
		fmt.Printf("    %-15s %15s\n", fmt.Sprintf("--%s", f.Name()), Yellow(f.Description()))
	}
}

func (cn *CommandNode) Run(args []string) {
	for _, f := range cn.flags {
		if err := f.Parse(); err != nil {
			fmt.Println(err)
			cn.Help()
			return
		}
		if f.Required() && !f.IsSet() {
			fmt.Printf(Red("\nRequired flag missing: ")+"%s\n", f.Name())
			cn.Help()
			return
		}
	}
	cn.RunFunc()
}

func (cn *CommandNode) NewBoolFlag(name, description string, required bool) {
	f := flags.NewBoolFlag(name, description, required)
	cn.flags = append(cn.flags, f)
	return
}

func (cn *CommandNode) NewIntFlag(name, description string, required bool) {
	f := flags.NewIntFlag(name, description, required)
	cn.flags = append(cn.flags, f)
	return
}

func (cn *CommandNode) NewStringFlag(name, description string, required bool) {
	f := flags.NewStringFlag(name, description, required)
	cn.flags = append(cn.flags, f)
	return
}
