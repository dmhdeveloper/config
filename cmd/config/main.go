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
	debug      *bool
)

var log = config.StandardLogger{}

var helpMessage = `
Config CLI is a wrapper for git.
To see git help commands, use git client directly eg. 'git -h'

To initialise the CLI, run 'config init' with the flags shown below:
`

func main() {
	initFlags := flag.NewFlagSet("init", flag.ExitOnError)
	repository = initFlags.String("git.url", "", "The git SSH url where your config and dotfiles are stored")
	gitDir = initFlags.String("git.dir", "~/.dotfiles", "The location to initialise the git bare repo")
	workTree = initFlags.String("work.tree", "~/", "The root location under which all files that you would like to save in git should be stored")
	sshKey = initFlags.String("ssh.key", "~/.ssh/id_rsa", "The SSH key used to interact with your private git repository storing your system config")
	debug = initFlags.Bool("debug", false, "Output all errors as they are returned from the CLI to stdout")

	if len(os.Args) == 1 {
		initFlags.PrintDefaults()
		os.Exit(0)
	}

	system := config.System{}
	scm := config.NewGit(system)
	cli := config.NewCLI(*debug, log, system, scm)

	switch os.Args[1] {
	case "help":
		fmt.Println(helpMessage)
		initFlags.PrintDefaults()
		os.Exit(0)
	case "version", "v":
		log.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		os.Exit(0)
	case "display":
		log.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		conf := cli.LoadGitConfig()

		contents, _ := yaml.Marshal(&conf)
		log.Println(string(contents))
	case "init":
		err := initFlags.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		conf := config.CLIConfig{
			GitDir:   *gitDir,
			WorkTree: *workTree,
		}

		file := fmt.Sprint(config.ConfigDir, "/", config.GitConfig)
		if ok := cli.UpdateConfigFile(file, conf); !ok {
			log.Println("failed to update config file")
			os.Exit(1)
		}

		if ok := cli.Init(*repository, conf.GitDir, *sshKey); !ok {
			log.Println("failed to initialise config cli")
			os.Exit(1)
		}
	default:
		conf := cli.LoadGitConfig()

		if conf.GitDir == "" || conf.WorkTree == "" {
			fmt.Println("To initialise the CLI, run `config init` with the flags shown below:")
			initFlags.PrintDefaults()
			os.Exit(1)
		}

		cli.Run(conf.GitDir, conf.WorkTree, os.Args[1:]...)
	}
}
