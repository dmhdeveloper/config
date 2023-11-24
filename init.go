package main

import (
	"bytes"
	"flag"
)

type initCmd struct {
	flags *flag.FlagSet
}

func NewInitCmd() initCmd {
	initFlags := flag.NewFlagSet("init", flag.ExitOnError)
	initFlags.String("url", "", "The git remote repository URL that stores your system configuration files.")
	initFlags.String("git.dir", "~/.dotfiles", "The git bare directory location.")
	initFlags.String("work.tree", "~/", "The root directory that git is allowed to view. All system config files should be discoverable within this root directory.")
	initFlags.String("ssh.key", "~/.ssh/id_rsa", "The ssh key used to interact with your git respository storing your configuration files.")
	return initCmd{
		flags: initFlags,
	}
}

func (i initCmd) Run(args ...string) int {
	err := i.flags.Parse(args)
	if err != nil {
		return 1
	}
	return 0
}

func (i initCmd) Help() string {
	buf := bytes.NewBuffer(make([]byte, 0))
	i.flags.SetOutput(buf)
	i.flags.PrintDefaults()
	return buf.String()
}
