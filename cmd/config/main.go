package main

import (
	"fmt"
	"os"

	"github.com/dmhdeveloper/config"
)

var (
	GitHash   string
	BuildTime string
	Version   string
)

func main() {
	if config.ShouldRunHelp() {
		config.RunHelp()
		os.Exit(1)
	}

	if config.ShouldInit() {
		err := config.RunInit()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	conf, err := config.LoadConfigFile(config.ConfigFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if config.IsEmpty(conf) {
		fmt.Println("ERROR: Config has not been initialised")
		fmt.Println("Run 'config init' to initialise the CLI")
		os.Exit(1)
	}

	if config.ShouldInit() {
		err = config.InitBareRepo(conf.GitDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = config.ConfigureRemoteForBareRepo(conf.SSHKey, conf.Repository, conf.GitDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if err := config.Initialised(conf); err != nil {
		fmt.Printf("ERROR: Config has encountered an error: %s\n", err)
		fmt.Println("Run 'config -h' or 'config --help' for more information about configuring the CLI")
		os.Exit(1)
	}

	err = config.ExecuteCommand(conf.WorkTree, conf.GitDir, os.Args[1:]...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
