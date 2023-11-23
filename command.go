package main

type Command interface {
	Run(args ...string) int
	Help() string
}
