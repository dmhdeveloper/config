package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dmhdeveloper/config"
	"gopkg.in/yaml.v3"
)

var (
	GitHash   string
	BuildTime string
	Version   string
)

var (
	repository *string
	gitDir     *string
	workTree   *string
	sshKey     *string
)

var helpMessage = `
Config CLI is a wrapper for git.
To see git help commands, use git client directly eg. 'git -h'

To initialise the CLI, run 'config init' with the flags shown below:
`

func main() {
	initFlags := flag.NewFlagSet("init", flag.PanicOnError)
	repository = initFlags.String("git.url", "", "The git SSH url where your config and dotfiles are stored")
	gitDir = initFlags.String("git.dir", "~/.dotfiles", "The location to initialise the git bare repo")
	workTree = initFlags.String("work.tree", "~/", "The root location under which all files that you would like to save in git should be stored")
	sshKey = initFlags.String("ssh.key", "~/.ssh/id_rsa", "The SSH key used to interact with your private git repository storing your system config")

	if len(os.Args) == 0 {
		initFlags.PrintDefaults()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println(helpMessage)
		initFlags.PrintDefaults()
		os.Exit(0)
	case "version", "v":
		fmt.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		os.Exit(0)
	case "display":
		fmt.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		conf, err := config.LoadConfigFile(config.FileLocation)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		contents, err := yaml.Marshal(&conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(contents))
		os.Exit(0)
	case "init":
		err := initFlags.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		conf := config.CLIConfig{
			CLI: config.CLI{
				GitDir:   *gitDir,
				WorkTree: *workTree,
			},
		}

		if config.IsEmpty(conf) {
			fmt.Println("To initialise the CLI, run `config init` with the flags shown below:")
			initFlags.PrintDefaults()
			os.Exit(1)
		}

		conf, err = config.UpdateConfigFile(config.FileLocation, conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = config.InitBareRepo(conf.CLI.GitDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = config.ConfigureRemoteForBareRepo(*sshKey, *repository, conf.CLI.GitDir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		conf, err := config.LoadConfigFile(config.FileLocation)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if config.IsEmpty(conf) {
			fmt.Println("To initialise the CLI, run `config init` with the flags shown below:")
			initFlags.PrintDefaults()
			os.Exit(1)
		}

		err = config.ExecuteCommand(conf.CLI.WorkTree, conf.CLI.GitDir, os.Args[1:]...)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
