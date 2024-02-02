package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type GitCmd struct {
	gitDir   string
	workTree string
	writer   io.Writer
}

func NewGitCmd(
	gitDir string,
	workTree string,
	writer io.Writer,
) GitCmd {
	return GitCmd{
		gitDir:   gitDir,
		workTree: workTree,
		writer:   writer,
	}
}

func (g GitCmd) Run(args ...string) int {
	defaults := []string{"--git-dir", g.gitDir, "--work-tree", g.workTree}
	defaults = append(defaults, args...)
	cmd := exec.Command("git", defaults...)
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

func (g GitCmd) Help() string {
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
	return fmt.Sprint(
		"This is a git wrapper, besides the list of commands above, all other flags are treated as git flags and passed to git.",
		"\n",
		"For git commands, see below: ",
		"\n\n",
		strings.Replace(buf.String(), "[-h | --help]", "--help", 1),
	)
}
