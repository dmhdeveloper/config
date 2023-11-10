package config

import (
	"errors"
	"io"
	"os"
	"os/user"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	FileLocation = "~/.config/config-cli/config.yml"
)

// This is a config file that you create on initiliasation.
// It stores sensitive data so it is not persisted in a remote repository.
// When you initialise config, it will create this in `~/.config/user-config/config.properties`
type CLIConfig struct {
	CLI CLI `yaml:"cli"`
}

type CLI struct {
	GitDir   string `yaml:"gitDir"`
	WorkTree string `yaml:"workTree"`
}

func LoadConfigFile(fileName string) (CLIConfig, error) {
	var conf CLIConfig
	c, err := openConfig(fileName, os.O_RDONLY)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(expandHomeDir("~/.config/config-cli"), 0777)
		if err != nil {
			return conf, err
		}

		_, err = os.Create(expandHomeDir(fileName))
		if err != nil {
			return conf, err
		}

		c, _ = openConfig(fileName, os.O_RDONLY)
	}
	if err != nil {
		return conf, err
	}

	contents, err := io.ReadAll(c)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(contents, &conf)
	return conf, nil
}

func UpdateConfigFile(fileName string, conf CLIConfig) (CLIConfig, error) {
	current, err := LoadConfigFile(fileName)
	if err != nil {
		return conf, err
	}

	if IsEqual(current, conf) {
		return conf, nil
	}

	c, err := openConfig(fileName, os.O_WRONLY|os.O_TRUNC)
	if err != nil {
		return conf, err
	}
	defer c.Close()

	err = c.Truncate(0)
	if err != nil {
		return conf, err
	}

	out, err := yaml.Marshal(&conf)
	if err != nil {
		return conf, err
	}

	_, err = c.Write(out)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func openConfig(fileName string, permissions int) (*os.File, error) {
	return os.OpenFile(expandHomeDir(fileName), permissions, 0666)
}

func IsEqual(src CLIConfig, dst CLIConfig) bool {
	return src.CLI.GitDir == dst.CLI.GitDir &&
		src.CLI.WorkTree == dst.CLI.WorkTree
}

func IsEmpty(conf CLIConfig) bool {
	return strings.TrimSpace(conf.CLI.GitDir) == "" ||
		strings.TrimSpace(conf.CLI.WorkTree) == ""
}

func (u CLIConfig) String() string {
	out, err := yaml.Marshal(&u)
	if err != nil {
		return ""
	}
	return string(out)
}

func expandHomeDir(s string) string {
	usr, err := user.Current()
	if err != nil {
		return s
	}

	dir := usr.HomeDir
	return strings.Replace(s, "~", dir, 1)
}
