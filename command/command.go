package command

type Command interface {
	Run(args ...string) int
	Help() string
}
