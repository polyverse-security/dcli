# DCLI

A library meant to facilitate cli argument parsing.

DiscoveryNode Interface:
````go
type DiscoveryNode interface {
	Run([]string) error
	Help()
	Name() string
	Description() string
}
````

Example usage:
```go
package mypackage
import (
	"github.com/polyverse-security/dcli"
	"fmt"
	)

var top = dcli.NewMenuNode(dcli.MenuNodeInput{
		Name:        "example application",
		Description: "example description",
	})
var subCommand = &dcli.CommandNode{
	N:       "subcommand",
	D:       "example subcommand",
	RunFunc: func(){
		fmt.Println("example run stuffs")
	},
}
subCommand.NewStringFlag(
		"arg1",
		"example arg 1 for subcommand",
		true,
	)
subCommand.NewStringFlag(
		"arg2",
		"example arg 2 for subcommand",
		false, 
	)
top.AddCommand(subCommand)
dcli.Start(top)
```