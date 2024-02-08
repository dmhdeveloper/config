package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dmhdeveloper/config/pkg/configs"
)

func RunGit(
	args []string,
) int {
	conf, err := configs.LoadGitConfig()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	defaults := []string{"--git-dir", conf.GitDir, "--work-tree", conf.WorkTree}
	defaults = append(defaults, args...)
	cmd := exec.Command("git", defaults...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
