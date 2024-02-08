package configs

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func read(filepath string, v interface{}) error {
	fi, err := os.Open(filepath)
	if err != nil {
		return err
	}

	contents, err := io.ReadAll(fi)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(contents, v)
	if err != nil {
		return err
	}

	return nil
}

func write(filepath string, v interface{}) error {
	content, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, content, 0644)
}
