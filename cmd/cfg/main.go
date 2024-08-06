package main

import (
	"fmt"
	"os"

	"github.com/dmhdeveloper/cfg/pkg/cli"
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

	os.Exit(cli.RunConfig(os.Args))
}
