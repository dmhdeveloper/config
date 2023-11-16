package config

import "fmt"

type Logger interface {
	Println(v ...any)
}

type StandardLogger struct{}

func (s StandardLogger) Println(v ...any) {
	fmt.Println(v...)
}
