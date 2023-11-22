package config

import (
	"fmt"
	"os"

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
	scm   SCM
}

func NewCLI(debug bool, log Logger, sys System, scm SCM) *CLI {
	return &CLI{
		debug: debug,
		log:   log,
		sys:   sys,
		scm:   scm,
	}
}

func (c CLI) LoadGitConfig() CLIConfig {
	var conf CLIConfig

	file := fmt.Sprint(ConfigDir, "/", GitConfig)
	exists, err := c.sys.FileExists(expandHomeDir(file))
	if err != nil || !exists {
		if c.debug {
			c.log.Println(err)
		}
		return conf
	}

	content, err := c.sys.GetFileContent(expandHomeDir(file))
	if err != nil || len(content) == 0 {
		if c.debug {
			c.log.Println(err)
		}
		return conf
	}

	_ = yaml.Unmarshal(content, &conf)
	return conf
}

func (c CLI) UpdateConfigFile(fileName string, conf CLIConfig) bool {
	var current CLIConfig
	path := expandHomeDir(fileName)

	content, err := c.sys.GetFileContent(path)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		return false
	}

	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		return false
	}

	if IsEqual(current, conf) {
		return true
	}

	out, err := yaml.Marshal(&conf)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		return false
	}

	err = c.sys.WriteFileContent(path, out)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		return false
	}
	return true
}

func (c CLI) Init(url string, scmDir string, sshKey string) bool {
	err := c.scm.Init(url, scmDir)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		c.log.Println("Configure source directory:\tfailed")
		return false
	}
	c.log.Println("Configure source directory:\tok")

	err = c.scm.ConfigureSSHKey(scmDir, sshKey)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		c.log.Println("Configure ssh key:\tfailed")
		return false
	}
	c.log.Println("Configure ssh key:\tok")

	err = c.scm.IgnoreUnknownFiles(scmDir, true)
	if err != nil {
		if c.debug {
			c.log.Println(err)
		}
		c.log.Println("Configure ignore unknown files:\tfailed")
		return false
	}
	c.log.Println("Configure ignore unknown files:\tok")
	return true
}

func (c CLI) Run(gitDir string, workTree string, commands ...string) {
	params := make([]string, 0)
	params = append(params, "--git-dir", gitDir)
	params = append(params, "--work-tree", workTree)
	params = append(params, os.Args[1:]...)
	c.sys.Run(
		"git",
		params...,
	)
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
