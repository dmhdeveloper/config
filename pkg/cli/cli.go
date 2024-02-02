package cli

import (
	"fmt"

	"github.com/dmhdeveloper/config/pkg/command"
)

const helpMessage = `
cfg - a wrapper for git to manage system config files.

usage: cfg [subcommand] <flags>

subcommands:
- help:       print help messsage.
- init:       initialize cfg cli.
- completion: output completion script.
- git:        interact with git to manage config.
       
example: 
    init
        cfg init -url <GIT_URL> -git.dir ~/.dotfiles -work.tree ~/ -ssh.key ~/.ssh/id_rsa
    git
        cfg git log --graph --oneline --abbrev-commit 
		completion
		    cfg completion zsh

completion:
    To enable completion, add the following to .zshrc
		eval $(cfg completion zsh)
`

func RunConfig(
	args []string,
) int {
	// No arguments passed, return help message
	if len(args) == 0 {
		fmt.Print(helpMessage)
		return 0
	}

	switch args[1] {
	case "help", "-h", "--help":
		fmt.Print(helpMessage)
	case "init":
		return command.RunInit(args[2:])
	case "git":
		return command.RunGit(args[2:])
	case "completion":
		return command.RunCompletion(args[2:])
	default:
		fmt.Printf("Uknown command: %s\n\n", args[1])
		fmt.Print(helpMessage)
	}

	// No error
	return 0
}
