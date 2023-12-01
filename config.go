package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	configDir     = ".config/config-cli"
	gitConfigFile = "git.yaml"
)

type CLIConfig struct {
	GitDir   string `yaml:"gitDir"`
	WorkTree string `yaml:"workTree"`
}

func LoadConfig() (CLIConfig, error) {
	var conf CLIConfig
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

		fi, err := os.Create(fullPath)
		if err != nil {
			return conf, err
		}

		contents, err := yaml.Marshal(conf)
		if err != nil {
			return conf, err
		}

		_, err = fi.Write(contents)
		if err != nil {
			return conf, err
		}

		err = fi.Close()
		if err != nil {
			return conf, err
		}
	}

	fi, err := os.Open(fullPath)
	if err != nil {
		return conf, err
	}

	contents, err := io.ReadAll(fi)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(contents, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func UpdateConfig(conf CLIConfig) (CLIConfig, error) {
	current, err := LoadConfig()
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

	content, err := yaml.Marshal(current)
	if err != nil {
		return conf, err
	}

	usr, err := user.Current()
	if err != nil {
		return conf, err
	}

	fullPath := fmt.Sprint(usr.HomeDir, "/", configDir, "/", gitConfigFile)
	err = os.WriteFile(fullPath, content, 0644)
	return current, err
}
