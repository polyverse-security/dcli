package dcli

import (
	"os"
)

func Start(top *CommandNode) {
	top.Run(os.Args)
}
