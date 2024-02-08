package configs

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

const (
	configDir     = ".config/config-cli"
	gitConfigFile = "git.yaml"
)

type Git struct {
	GitDir   string `yaml:"gitDir"`
	WorkTree string `yaml:"workTree"`
}

func LoadGitConfig() (Git, error) {
	var conf Git
	usr, err := user.Current()
	if err != nil {
		return conf, err
	}

	fullPath := fmt.Sprint(usr.HomeDir, "/", configDir, "/", gitConfigFile)

	_, err = os.Stat(fullPath)
	if err != nil {
		if err := os.MkdirAll(fmt.Sprint(usr.HomeDir, "/", configDir), 0777); err != nil {
			return conf, err
		}

		_, err := os.Create(fullPath)
		if err != nil {
			return conf, err
		}

		err = write(fullPath, &conf)
		if err != nil {
			return conf, err
		}
	}

	err = read(fullPath, &conf)
	return conf, err
}

func UpdateGitConfig(conf Git) (Git, error) {
	current, err := LoadGitConfig()
	if err != nil {
		return conf, err
	}

	// Guard against overwriting usuable values with empty values
	if strings.TrimSpace(conf.GitDir) != "" {
		current.GitDir = conf.GitDir
	}
	if strings.TrimSpace(conf.WorkTree) != "" {
		current.WorkTree = conf.WorkTree
	}

	usr, err := user.Current()
	if err != nil {
		return conf, err
	}

	fullPath := fmt.Sprint(usr.HomeDir, "/", configDir, "/", gitConfigFile)
	err = write(fullPath, &current)
	return current, err
}
