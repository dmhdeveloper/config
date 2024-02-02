package command

import (
	_ "embed"
	"fmt"
)

const (
	bashBaseDir = "/usr/share/bash-completion/completions/"
)

//go:embed completion/cfg.zsh
var zshCompletionScript string

//go:embed completion/cfg.bash
var bashCompletionScript string

func RunCompletion(args []string) int {
	if len(args) == 0 {
		return 1
	}

	switch args[0] {
	case "zsh":
		fmt.Print(zshCompletionScript)
	case "bash":
		fmt.Print(bashCompletionScript)
	}
	return 0
}
