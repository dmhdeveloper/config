package config

import (
	"fmt"
)

type SCM interface {
	Init(url string, dir string) bool
	ConfigureSSHKey(dir string, sshKeyPath string) bool
	IgnoreUnknownFiles(dir string, ignore bool) bool
}

type Git struct {
	debug bool
	log   Logger
	sys   OS
}

func NewGit(debug bool, log Logger, sys System) *Git {
	return &Git{
		debug: debug,
		log:   log,
		sys:   sys,
	}
}

func (g Git) Init(url string, dir string) bool {
	err := g.sys.Run("git", "init", "--bare", dir)
	if err != nil {
		if g.debug {
			g.log.Println(err)
		}
		return false
	}

	err = g.sys.Run("git", "--git-dir", dir, "remote", "add", "origin", url)
	if err != nil {
		if g.debug {
			g.log.Println(err)
		}
		return false
	}
	return true
}

func (g Git) ConfigureSSHKey(gitDir string, sshKeyPath string) bool {
	err := g.sys.Run("git", "--git-dir", gitDir, "config", "--local", "core.sshCommand", fmt.Sprintf("ssh -i %s", sshKeyPath))
	if err != nil {
		if g.debug {
			g.log.Println(err)
		}
		return false
	}
	return true
}

func (g Git) IgnoreUnknownFiles(gitDir string, ignore bool) bool {
	value := "no"
	if ignore {
		value = "yes"
	}
	err := g.sys.Run("git", "--git-dir", gitDir, "config", "--local", "status.showUntrackedFiles", value)
	if err != nil {
		if g.debug {
			g.log.Println(err)
		}
		return false
	}
	return true
}
