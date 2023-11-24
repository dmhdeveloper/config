package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

type gitCmd struct {
	gitDir   string
	workTree string
	writer   io.Writer
}

func NewGitCmd(
	gitDir string,
	workTree string,
	writer io.Writer,
) gitCmd {
	return gitCmd{
		gitDir:   gitDir,
		workTree: workTree,
		writer:   writer,
	}
}

func (g gitCmd) Run(args ...string) int {
	cmd := exec.Command("git", args...)
	cmd.Stdout = g.writer
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err.(*exec.ExitError).ExitCode()
	}

	err = cmd.Wait()
	if err != nil {
		return err.(*exec.ExitError).ExitCode()
	}
	return 0
}

func (g gitCmd) Help() string {
	buf := bytes.NewBuffer(make([]byte, 0))
	cmd := exec.Command("git", "-h")
	cmd.Stdout = buf

	err := cmd.Start()
	if err != nil {
		return err.Error()
	}

	err = cmd.Wait()
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
