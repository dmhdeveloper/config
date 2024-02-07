package command

import "fmt"

const (
	bashBaseDir = "/usr/share/bash-completion/completions/"
	zshBaseDir  = "/usr/share/zsh/site-functions/"
)

type CompletionCmd struct{}

func NewCompletionCmd() CompletionCmd {
	return CompletionCmd{}
}

func (c CompletionCmd) Run(args ...string) int {
	fmt.Println(args)
	return 0
}

func (c CompletionCmd) Help() string {
	return fmt.Sprint(bashBaseDir, zshBaseDir)
}
