package dcli

import (
	"fmt"
	"os"
	"strings"

	"github.com/polyverse-security/dcli/flags"
	"github.com/polyverse-security/dcli/toggles"
)

type SingleNode struct {
	D       string // Description
	N       string // Name
	flags   []flags.Flag
	toggles []*toggles.Toggle
}

func (n *SingleNode) NewBoolFlag(name, description string, required bool) {
	f := flags.NewBoolFlag(name, description, required)
	n.flags = append(n.flags, f)
	return
}

func (n *SingleNode) NewIntFlag(name, description string, required bool) {
	f := flags.NewIntFlag(name, description, required)
	n.flags = append(n.flags, f)
	return
}

func (n *SingleNode) NewStringFlag(name, description string, required bool) {
	f := flags.NewStringFlag(name, description, required)
	n.flags = append(n.flags, f)
	return
}

func (n *SingleNode) NewToggle(name, description string) {
	n.toggles = append(n.toggles, toggles.NewToggle(name, description))
	return
}

// Parse handles the parsing of flags and toggles
func (n SingleNode) Parse() error {
	UsageSlice = append(UsageSlice, n.N)
	// Parse out the toggles
	togglesRemovedArgs := toggles.ParseToggles(os.Args[1:])
	// Parse out the flags
	flagsRemovedArgs := flags.ParseFlags(togglesRemovedArgs)
	_ = flagsRemovedArgs

	for _, f := range n.flags {
		if err := f.Parse(); err != nil {
			fmt.Println(err)
			return err
		}
		if f.Required() && !f.IsSet() {
			fmt.Printf(Red("\nRequired flag missing: ")+"%s\n", f.Name())
			n.Help()
			return fmt.Errorf("missing required flag %s", f.Name())
		}
	}
	return nil
}

func (n *SingleNode) Help() {
	fmt.Println()
	n.printUsage()
	n.printDescription()
	n.printFlags()
	n.printToggles()
	fmt.Println()
}

func (n *SingleNode) printUsage() {
	fmt.Println(Cyan(fmt.Sprintf("\nUsage:")))
	fmt.Printf("    %-15s [flags]\n", strings.Join(UsageSlice, " "))
}

func (n *SingleNode) printFlags() {
	fmt.Println(Cyan(fmt.Sprintf("\nFlags:")))
	for _, f := range n.flags {
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

func (n *SingleNode) printDescription() {
	fmt.Println(Cyan(fmt.Sprintf("\nDescription:")))
	fmt.Printf("    %-15s\n", n.D)
}

func (n *SingleNode) printToggles() {
	if len(n.toggles) == 0 {
		return
	}

	fmt.Println(Cyan(fmt.Sprintf("\nToggles:")))
	for _, t := range n.toggles {
		fmt.Printf("    %-15s %15s\n",
			fmt.Sprintf("--%s", t.Name()),
			Yellow(t.Description()),
		)
	}
}
