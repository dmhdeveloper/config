package main

import (
	"fmt"
	"os"

	"github.com/dmhdeveloper/config/command"
	"github.com/dmhdeveloper/config/configs"
	"github.com/dmhdeveloper/config/logger"
)

var (
	GitHash   string
	BuildTime string
	Version   string
)

var log = logger.FmtLogger{}

func main() {
	if len(os.Args) == 1 {
		log.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		os.Exit(0)
	}

	_, err := configs.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		command := command.NewInitCmd(log)
		os.Exit(command.Run(os.Args[2:]...))
	default:
		log.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
	}
}
