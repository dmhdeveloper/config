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

	conf, err := configs.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		command := command.NewInitCmd(log)
		os.Exit(command.Run(os.Args[2:]...))
	case "-h":
		helpMessage := command.RunHelp(command.InitCmd{}.Help, command.GitCmd{}.Help)
		log.Println(helpMessage)
	default:
		command := command.NewGitCmd(conf.GitDir, conf.WorkTree, os.Stdout)
		os.Exit(command.Run(os.Args[1:]...))
	}
}
