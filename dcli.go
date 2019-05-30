package dcli

import (
	"os"

	"github.com/fatih/color"
	"github.com/polyverse-security/dcli/flags"
	"github.com/polyverse-security/dcli/toggles"
)

var (
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	Pink   = color.New(color.FgHiMagenta).SprintFunc()
)

type DiscoveryNode interface {
	Run([]string) error
	Help()
	Name() string
	Description() string
}

// UsageSlice contains the Name of each node that has been called. This helps build the Usage: help output
var UsageSlice []string

// Start begins parsing the os.Args and executes the nodes.
func Start(top DiscoveryNode) error {
	// Parse out the toggles
	togglesRemovedArgs := toggles.ParseToggles(os.Args[1:])
	// Parse out the flags
	flagsRemovedArgs := flags.ParseFlags(togglesRemovedArgs)
	// Begin running nodes
	return top.Run(flagsRemovedArgs)
}

// New returns a MenuNode node intended to be the top level MenuNode
func New(serviceName, description string) *MenuNode {
	var top = &MenuNode{
		N: serviceName,
		D: description,
	}
	return top
}
