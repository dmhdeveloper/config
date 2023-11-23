package main

import "fmt"

type Logger interface {
	Println(v ...any)
}

type FmtLogger struct{}

func (f FmtLogger) Println(v ...any) {
	fmt.Println(v...)
}
