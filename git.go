package config

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func InitBareRepo(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return errors.New("directory already exists, aborting")
	}
	return runGit("init", "--bare", path)
}

func ConfigureRemoteForBareRepo(sshKeyPath string, remoteURL string, gitDir string) error {
	_, err := os.Stat(gitDir)
	switch err != nil {
	case errors.Is(err, os.ErrExist):
	default:
		return errors.New("directory does not exist, aborting")
	}
	if err := runGit("--git-dir", gitDir, "remote", "add", "origin", remoteURL); err != nil {
		return err
	}
	if err := runGit("--git-dir", gitDir, "config", "--local", "status.showUntrackedFiles", "no"); err != nil {
		return err
	}
	if err := runGit("--git-dir", gitDir, "config", "--local", "core.sshCommand", fmt.Sprintf("ssh -i %s", sshKeyPath)); err != nil {
		return err
	}
	return runGit("--git-dir", gitDir, "fetch", "-p")
}

func ExecuteCommand(workDir string, gitDir string, extraOptions ...string) error {
	options := []string{
		"--git-dir", gitDir,
		"--work-tree", workDir,
	}
	options = append(options, extraOptions...)
	return runGit(options...)
}

func runGit(commands ...string) error {
	cmd := exec.Command("git", commands...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Wait()
}
