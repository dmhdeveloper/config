package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
)

const (
	ConfigFile = ".config/user-config/config.properties"
)

const (
	RepositoryKey = "repository"
	GitDirKey     = "git.dir"
	WorkTreeKey   = "work.tree"
	SSHKey        = "ssh.key"
)

const (
	defaultWorkTree = "~/"
	defaultGitDir   = "~/.dotfiles"
	defaultSSHKey   = "~/.ssh/id_rsa"
)

// This is a config file that you create on initiliasation.
// It stores sensitive data so it is not persisted in a remote repository.
// When you initialise config, it will create this in `~/.config/user-config/config.properties`
type UserConfig struct {
	Repository string
	GitDir     string
	WorkTree   string
	SSHKey     string
}

func LoadConfigFile(fileName string) (UserConfig, error) {
	var conf UserConfig
	c, err := openConfig(fileName, os.O_RDONLY)
	if errors.Is(err, os.ErrNotExist) {
		usr, err := user.Current()
		if err != nil {
			return conf, err
		}

		dir := usr.HomeDir
		err = os.MkdirAll(fmt.Sprintf("%s/%s", dir, ".config/user-config"), 0777)
		if err != nil {
			return conf, err
		}

		_, err = os.Create(fmt.Sprintf("%s/%s", dir, fileName))
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

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		switch {
		case strings.TrimSpace(key) == RepositoryKey:
			conf.Repository = strings.TrimSpace(value)
		case strings.TrimSpace(key) == GitDirKey:
			conf.GitDir = strings.TrimSpace(value)
		case strings.TrimSpace(key) == WorkTreeKey:
			conf.WorkTree = strings.TrimSpace(value)
		case strings.TrimSpace(key) == SSHKey:
			conf.SSHKey = strings.TrimSpace(value)
		}
	}

	return conf, nil
}

func UpdateConfigFile(fileName string, conf UserConfig) (UserConfig, error) {
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

	err = c.Truncate(0)
	if err != nil {
		return conf, err
	}

	err = writeConfig(c, conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func openConfig(fileName string, permissions int) (*os.File, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	dir := usr.HomeDir
	return os.OpenFile(fmt.Sprintf("%s/%s", dir, fileName), permissions, 0666)
}

func writeConfig(writer io.Writer, conf UserConfig) error {
	contents := make([]byte, 0)
	buf := bytes.NewBuffer(contents)
	buf.WriteString(fmt.Sprintf("%s = %s\n", RepositoryKey, conf.Repository))
	buf.WriteString(fmt.Sprintf("%s = %s\n", GitDirKey, conf.GitDir))
	buf.WriteString(fmt.Sprintf("%s = %s\n", WorkTreeKey, conf.WorkTree))
	buf.WriteString(fmt.Sprintf("%s = %s\n", SSHKey, conf.SSHKey))

	_, err := writer.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func IsEqual(src UserConfig, dst UserConfig) bool {
	return src.Repository == dst.Repository &&
		src.GitDir == dst.GitDir &&
		src.WorkTree == dst.WorkTree &&
		src.SSHKey == dst.SSHKey
}

func IsEmpty(conf UserConfig) bool {
	return strings.TrimSpace(conf.Repository) == "" ||
		strings.TrimSpace(conf.GitDir) == "" ||
		strings.TrimSpace(conf.WorkTree) == "" ||
		strings.TrimSpace(conf.SSHKey) == ""
}

func (u UserConfig) String() string {
	var s string
	s = fmt.Sprintf("%s = %s\n", RepositoryKey, u.Repository)
	s = fmt.Sprintf("%s%s = %s\n", s, GitDirKey, u.GitDir)
	s = fmt.Sprintf("%s%s = %s\n", s, WorkTreeKey, u.WorkTree)
	s = fmt.Sprintf("%s%s = %s", s, SSHKey, u.SSHKey)
	return s
}
