package config

import (
	"fmt"
)

type SCM interface {
	Init(url string, dir string) error
	ConfigureSSHKey(dir string, sshKeyPath string) error
	IgnoreUnknownFiles(dir string, ignore bool) error
}

type Git struct {
	sys System
}

func NewGit(sys System) *Git {
	return &Git{
		sys: sys,
	}
}

func (g Git) Init(url string, dir string) error {
	err := g.sys.Run("git", "init", "--bare", dir)
	if err != nil {
		return err
	}

	return g.sys.Run("git", "--git-dir", dir, "remote", "add", "origin", url)
}

func (g Git) ConfigureSSHKey(gitDir string, sshKeyPath string) error {
	return g.sys.Run("git", "--git-dir", gitDir, "config", "--local", "core.sshCommand", fmt.Sprint("ssh -i ", sshKeyPath))
}

func (g Git) IgnoreUnknownFiles(gitDir string, ignore bool) error {
	value := "no"
	if ignore {
		value = "yes"
	}
	return g.sys.Run("git", "--git-dir", gitDir, "config", "--local", "status.showUntrackedFiles", value)
}
