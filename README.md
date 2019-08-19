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
package main

import (
	"fmt"

	"github.com/polyverse-security/dcli"
	"github.com/polyverse-security/dcli/flags"
)

func main() {
	var top = dcli.NewMenuNode(dcli.MenuNodeInput{
		Name:        "scratchpad",
		Description: "example implementation of dcli",
	})

	var boolCMD = &dcli.CommandNode{
		N:       "test-bool",
		D:       "test the bool flag",
		RunFunc: testBool,
	}

	var intCMD = &dcli.CommandNode{
		N:       "test-int",
		D:       "test the int flag",
		RunFunc: testInt,
	}

	var strCMD = &dcli.CommandNode{
		N:       "test-string",
		D:       "test the bool flag",
		RunFunc: testStr,
	}

	top.AddSubCommand(boolCMD)
	boolCMD.NewBoolFlag("required", "a required flag", true)
	boolCMD.NewBoolFlag("not-required", "not a required flag", false)

	top.AddSubCommand(intCMD)
	intCMD.NewIntFlag("required", "a required flag", true)
	intCMD.NewIntFlag("not-required", "not a required flag", false)

	top.AddSubCommand(strCMD)
	strCMD.NewStringFlag("required", "a required flag", true)
	strCMD.NewStringFlag("not-required", "not a required flag", false)

	// Start
	dcli.Start(top)
}

func testBool() {
	required := *flags.GetBool("required").Value()
	var notRequired bool
	if flags.GetBool("not-required").IsSet() {
		notRequired = *flags.GetBool("not-required").Value()
	}
	fmt.Println("Required:", required)
	fmt.Println("Not Required:", notRequired)
}

func testStr() {
	required := *flags.GetStringFlag("required").Value()
	var notRequired string
	if flags.GetStringFlag("not-required").IsSet() {
		notRequired = *flags.GetStringFlag("not-required").Value()
	}
	fmt.Println("Required:", required)
	fmt.Println("Not Required:", notRequired)
}

func testInt() {
	required := *flags.GetIntFlag("required").Value()
	var notRequired int
	if flags.GetIntFlag("not-required").IsSet() {
		notRequired = *flags.GetIntFlag("not-required").Value()
	}
	fmt.Println("Required:", required)
	fmt.Println("Not Required:", notRequired)
}

```
# SingleNode
SingleNode should be used when you don't have multiple routing options. It provides the flag and toggle checking and
 exposure like a CommandNode but without the MenuNode requirement. It does automatically run a registered function. 

Example SingleNode
```
func startup() (bbservice.Service, error) {
   	node := dcli.SingleNode{
   		N: "duality",
   		D: "BigBang service responsible for populating an Instance with unscrambled packages to fulfill the Closed Loop Instance requirements.",
   	}
   	node.NewStringFlag("account", "aws account id", true)
   	node.NewStringFlag("region", "aws region", true)
   	node.NewStringFlag("stack", "BigBang stack namespace", true)
   
   	if err := node.Parse(); err != nil {
   		return nil, err
   	}
   
   	return AwsDuality{}, nil
}
```


### Possibly TODO:
* Support POSIX arguments
* Support global arguments
* Tests
* More documentation
* Automatic argument to variable binding
* Better error handling/help printing