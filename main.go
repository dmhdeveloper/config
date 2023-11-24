package main

import (
	"fmt"
	"os"
)

var (
	GitHash   string
	BuildTime string
	Version   string
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		os.Exit(0)
	}

	conf, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(conf)
}
