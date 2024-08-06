package command

import (
	"bytes"
	"flag"

	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/dmhdeveloper/cfg/pkg/configs"
)

func RunInit(
	args []string,
) int {
	var (
		url      string
		gitDir   string
		workTree string
		sshKey   string
	)

	flags := flag.NewFlagSet("init", flag.ExitOnError)
	flags.StringVar(&url, "url", "", "(Required) The git remote repository URL that stores your system configuration files.")
	flags.StringVar(&gitDir, "git.dir", "(Required) ~/.dotfiles", "The git bare directory location.")
	flags.StringVar(&workTree, "work.tree", "~/", "(Required) All system config files should be discoverable within this root directory.")
	flags.StringVar(&sshKey, "ssh.key", "~/.ssh/id_rsa", "(Required) The ssh key used to interact with your git repository storing your configuration files.")

	err := flags.Parse(args)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	conf := configs.Git{
		GitDir:   gitDir,
		WorkTree: workTree,
	}
	_, err = configs.UpdateGitConfig(conf)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	if err := runGit(os.Stdout, "init", "--bare", gitDir); err != nil {
		fmt.Println(err)
		return 1
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	if err := runGit(buf, "--git-dir", gitDir, "remote", "show"); err != nil {
		fmt.Println(err)
		return 1
	}

	// We are re-initialising the repo, the origin might already be set
	if strings.TrimSpace(buf.String()) != "origin" {
		if err := runGit(os.Stdout, "--git-dir", gitDir, "remote", "add", "origin", url); err != nil {
			fmt.Println(err)
			return 1
		}
	} else {
		if err := runGit(os.Stdout, "--git-dir", gitDir, "remote", "set-url", "origin", url); err != nil {
			fmt.Println(err)
			return 1
		}
	}

	if err := runGit(os.Stdout, "--git-dir", gitDir, "config", "--local", "status.showUntrackedFiles", "no"); err != nil {
		fmt.Println(err)
		return 1
	}
	if err := runGit(os.Stdout, "--git-dir", gitDir, "config", "--local", "core.sshCommand", fmt.Sprintf("ssh -i %s", sshKey)); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func runGit(out io.Writer, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
