package command

import "fmt"

const (
	defaultHelpMessage = `usage: config -h

  Display this help message.
`
)

type Command interface {
	Run(args ...string) int
	Help() string
}

func RunHelp(help ...func() string) string {
	var helpMessage string
	for _, f := range help {
		helpMessage = fmt.Sprint(helpMessage, "\n", f())
	}
	return fmt.Sprint(defaultHelpMessage, helpMessage)
}
