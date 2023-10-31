package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"
)

const helpMessage = `Config is a git wrapper for managing rc and other config files on your system.
It uses a git bare repo to manage your config while adding a few additional flags to make the experience more pleasant.

The config that the CLI uses to know what remote repository to use and how to interact with it
is stored in the following location:

$ ~/.config/user-config/config.properties

It contains the following attributes:

- repository                       # The remote git ssh URL that will be used to store config files
- git.dir (default: ~/.dotfiles)   # The location of the git bare repo files
- work.tree (default: ~/)          # The root location that config should use when looking for files, default to $HOME
- ssh.key (default: ~/.ssh/id_rsa) # Config interacts with your repo using SSH. The path to the SSH key.

To configure the CLI service, run

$ config init

For standard git commands, see below:`

func ShouldRunHelp() bool {
	return len(os.Args) == 1 || (os.Args[1] == "-h" || os.Args[1] == "--help")
}

func RunHelp() {
	fmt.Println(helpMessage)
	fmt.Println("")
	err := runGit("-h")
	if err != nil {
		fmt.Println(err)
	}
}

func ShouldInit() bool {
	return len(os.Args) > 1 && os.Args[1] == "init"
}

func Initialised(conf UserConfig) error {
	_, err := os.Stat(conf.GitDir)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	_, err = os.Stat(conf.WorkTree)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	_, err = os.Stat(conf.SSHKey)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func RunInit() error {
	conf, err := LoadConfigFile(ConfigFile)
	if err != nil {
		return err
	}

	repo := inputPromptWithDefault("Repository", conf.Repository, "")
	if repo == "" {
		return errors.New("repository is mandatory for this CLI")
	}

	gitDir := inputPromptWithDefault("Git Bare Directory", conf.GitDir, defaultGitDir)
	gitDir, err = expandHomeDir(gitDir)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(gitDir, "/") {
		return fmt.Errorf("only absolute paths allowed: %s", gitDir)
	}

	workTree := inputPromptWithDefault("Work Tree", conf.WorkTree, defaultWorkTree)
	workTree, err = expandHomeDir(workTree)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(workTree, "/") {
		return fmt.Errorf("only absolute paths allowed: %s", workTree)
	}

	sshKey := inputPromptWithDefault("SSH Key", conf.SSHKey, defaultSSHKey)
	sshKey, err = expandHomeDir(sshKey)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(sshKey, "/") {
		return fmt.Errorf("only absolute paths allowed: %s", sshKey)
	}

	conf.Repository = repo
	conf.GitDir = gitDir
	conf.WorkTree = workTree
	conf.SSHKey = sshKey

	_, err = UpdateConfigFile(ConfigFile, conf)
	return err
}

// InputPrompt receives a string value using the label
func inputPromptWithDefault(label, current, deflt string) string {
	var s string
	r := bufio.NewReader(os.Stdin)

	var prompt string
	switch {
	case deflt == "" && current == "":
		prompt = fmt.Sprintf("%s: ", label)
	case current != "":
		prompt = fmt.Sprintf("%s (Current: %s): ", label, current)
	case deflt != "":
		prompt = fmt.Sprintf("%s (Default: %s): ", label, deflt)
	default:
		prompt = fmt.Sprintf("%s: ", label)
	}

	for {
		fmt.Fprint(os.Stderr, prompt)
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	input := strings.TrimSpace(s)
	input = expandEnvVars(input)
	if input == "" {
		input = current
	}
	if input == "" {
		input = deflt
	}
	return input
}

func expandEnvVars(s string) string {
	for strings.Contains(s, "$") {
		start := strings.Index(s, "$")
		end := start
		for end < len(s) {
			end++
			if s[end] == '/' {
				break
			}
		}
		envVar := s[start:end]
		env := strings.ReplaceAll(envVar, "$", "")
		// Support ${ENV}
		env = strings.ReplaceAll(env, "{", "")
		env = strings.ReplaceAll(env, "}", "")
		value := os.Getenv(env)
		s = strings.Replace(s, envVar, value, 1)
	}
	return s
}

func expandHomeDir(s string) (string, error) {
	if strings.HasPrefix(s, "~/") {
		usr, err := user.Current()
		if err != nil {
			return s, err
		}

		dir := usr.HomeDir
		s = strings.Replace(s, "~", dir, 1)
	}
	return s, nil
}
