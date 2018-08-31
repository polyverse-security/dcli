package dcli

import "fmt"

// CommandNode is a node that executes a function.
type CommandNode struct {
	D        string         // Description
	N        string         // Name
	U        string         // Usage
	Prompt   bool           // Require a prompt to execute the RunFunc
	RunFunc  func([]string) // the function to run when this node is called
	ArgCount int            // the number of args this node is required to receive
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
	fmt.Printf("Usage: %s\n\n", cn.U)
}

func (cn *CommandNode) Run(args []string) {
	if len(args) < cn.ArgCount {
		cn.Help()
		return
	}

	/*if cn.Prompt {
		fmt.Printf("Are you sure you want to do this? y/n:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			switch strings.ToLower(scanner.Text()) {
			case "y":
				break
			case "n":
				fmt.Println("\noperation cancelled by user input.")
				return
			default:
				continue
			}
		}

		if scanner.Err() != nil {
			fmt.Println("error while getting user input:", scanner.Err())
			return
		}
	}*/

	cn.RunFunc(args)
}
