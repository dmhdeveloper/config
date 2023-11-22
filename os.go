package config

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type OS interface {
	GetFileContent(file string) ([]byte, error)
	WriteFileContent(file string, content []byte) error
	DirExists(name string) (bool, error)
	FileExists(name string) (bool, error)
	Run(name string, commands ...string) error
	RunWithOutput(name string, commands ...string) (string, error)
}

type System struct{}

func (s System) DirExists(
	name string,
) (bool, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func (s System) FileExists(name string) (bool, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return false, err
	}

	return !fi.IsDir(), nil
}

func (s System) Run(name string, commands ...string) error {
	cmd := exec.Command(name, commands...)
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		return err
	}

	return cmd.Wait()
}

func (s System) RunWithOutput(name string, commands ...string) (string, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	cmd := exec.Command(name, commands...)
	cmd.Stdout = buf

	err := cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	return buf.String(), err
}

func (s System) GetFileContent(file string) ([]byte, error) {
	var content []byte

	exists, err := s.FileExists(file)
	if err != nil {
		return content, err
	}

	if !exists {
		_, err = os.Create(file)
		if err != nil {
			return content, err
		}
	}

	c, err := os.OpenFile(expandHomeDir(file), os.O_RDONLY, 0666)
	if err != nil {
		return content, err
	}

	return io.ReadAll(c)
}

func (s System) WriteFileContent(file string, content []byte) error {
	c, err := os.OpenFile(expandHomeDir(file), os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.Truncate(0)
	if err != nil {
		return err
	}

	_, err = c.Write(content)
	return err
}

func expandHomeDir(s string) string {
	usr, err := user.Current()
	if err != nil {
		return s
	}

	dir := usr.HomeDir
	return strings.Replace(s, "~", dir, 1)
}
