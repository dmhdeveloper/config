package command

import (
	_ "embed"
	"fmt"
)

//go:embed completion/cfg.zsh
var zshCompletionScript string

func RunCompletion(args []string) int {
	if len(args) == 0 {
		return 1
	}

	if args[0] == "zsh" {
		fmt.Print(zshCompletionScript)
	}
	return 0
}
