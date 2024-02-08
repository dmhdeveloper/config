package cli

import (
	"fmt"

	"github.com/dmhdeveloper/config/pkg/command"
)

const helpMessage = ""

func RunConfig(
	args []string,
) int {
	// No arguments passed, return help message
	if len(args) == 0 {
		fmt.Println(helpMessage)
		return 0
	}

	switch args[1] {
	case "help", "-h", "--help":
		fmt.Println(helpMessage)
	case "init":
		return command.RunInit(args[2:])
	case "git":
		return command.RunGit(args[2:])
	default:
		fmt.Printf("Uknown command: %s\n\n", args[1])
		fmt.Println(helpMessage)
	}

	// No error
	return 0
}
