// TODO when parsing flags for this node, place them in a global map where the CommandNode's RunFunc can access them
package dcli

import (
	"fmt"
	"strings"

	"github.com/polyverse-security/dcli/flags"
	"github.com/polyverse-security/dcli/toggles"
)

// CommandNode is a node that executes a function.
type CommandNode struct {
	D       string // Description
	N       string // Name
	Prompt  bool   // Require a prompt to execute the RunFunc
	RunFunc func() // the function to run when this node is called
	flags   []flags.Flag
	toggles []*toggles.Toggle
}

func (cn *CommandNode) Name() string {
	return cn.N
}

func (cn *CommandNode) Description() string {
	return cn.D
}

func (cn *CommandNode) printUsage() {
	fmt.Println(Cyan(fmt.Sprintf("\nUsage:")))
	fmt.Printf("    %-15s [flags]\n", strings.Join(UsageSlice, " "))
}

func (cn *CommandNode) printFlags() {
	fmt.Println(Cyan(fmt.Sprintf("\nFlags:")))
	for _, f := range cn.flags {
		switch f.Required() {
		case true:
			fmt.Printf("    %-15s %15s %s\n",
				fmt.Sprintf("--%s", f.Name()),
				Yellow(f.Description()),
				Red("(required)"),
			)
		case false:
			fmt.Printf("    %-15s %15s\n",
				fmt.Sprintf("--%s", f.Name()),
				Yellow(f.Description()),
			)
		}

	}
}

func (cn *CommandNode) printDescription() {
	fmt.Println(Cyan(fmt.Sprintf("\nDescription:")))
	fmt.Printf("    %-15s\n", cn.D)
}

func (cn *CommandNode) printToggles() {
	if len(cn.toggles) == 0 {
		return
	}

	fmt.Println(Cyan(fmt.Sprintf("\nToggles:")))
	for _, t := range cn.toggles {
		fmt.Printf("    %-15s %15s\n",
			fmt.Sprintf("--%s", t.Name()),
			Yellow(t.Description()),
		)
	}
}

func (cn *CommandNode) Help() {
	fmt.Println()
	cn.printUsage()
	cn.printDescription()
	cn.printFlags()
	cn.printToggles()
	fmt.Println()
}

func (cn *CommandNode) Run(args []string) error {
	UsageSlice = append(UsageSlice, cn.N)

	for _, f := range cn.flags {
		if err := f.Parse(); err != nil {
			fmt.Println(err)
			return err
		}
		if f.Required() && !f.IsSet() {
			fmt.Printf(Red("\nRequired flag missing: ")+"%s\n", f.Name())
			cn.Help()
			return fmt.Errorf("missing required flag %s", f.Name())
		}
	}

	cn.RunFunc()
	return nil
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

func (cn *CommandNode) NewToggle(name, description string) {
	cn.toggles = append(cn.toggles, toggles.NewToggle(name, description))
	return
}
