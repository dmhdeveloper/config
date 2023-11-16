package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	ConfigDir = "~/.config/config-cli"
	GitConfig = "git.yml"
)

type CLIConfig struct {
	GitDir   string `yaml:"gitDir"`
	WorkTree string `yaml:"workTree"`
}

type CLI struct {
	debug bool
	log   Logger
	sys   System
}

func NewCLI(debug bool, log Logger, sys System) *CLI {
	return &CLI{
		debug: debug,
		log:   log,
		sys:   sys,
	}
}

func (c CLI) LoadGitConfig() CLIConfig {
	var conf CLIConfig
	exists, err := c.sys.FileExists(fmt.Sprint(ConfigDir, "/", GitConfig))
	if err != nil || !exists {
		if c.debug {
			c.log.Println(err)
		}
		return conf
	}

	content, err := c.sys.GetFileContent(fmt.Sprint(ConfigDir, "/", GitConfig))
	if err != nil || len(content) == 0 {
		if c.debug {
			c.log.Println(err)
		}
		return conf
	}

	_ = yaml.Unmarshal(content, &conf)
	return conf
}

func (c CLI) UpdateConfigFile(fileName string, conf CLIConfig) {
	var current CLIConfig

	content, err := c.sys.GetFileContent(fileName)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		return
	}

	_ = yaml.Unmarshal(content, &conf)

	if IsEqual(current, conf) {
		return
	}

	out, err := yaml.Marshal(&conf)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		return
	}

	err = c.sys.WriteFileContent(fileName, out)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
	}
}

func IsEqual(src CLIConfig, dst CLIConfig) bool {
	return src.GitDir == dst.GitDir &&
		src.WorkTree == dst.WorkTree
}

func (u CLIConfig) String() string {
	out, err := yaml.Marshal(&u)
	if err != nil {
		return ""
	}
	return string(out)
}
