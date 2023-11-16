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
	initFlags := flag.NewFlagSet("init", flag.PanicOnError)
	repository = initFlags.String("git.url", "", "The git SSH url where your config and dotfiles are stored")
	gitDir = initFlags.String("git.dir", "~/.dotfiles", "The location to initialise the git bare repo")
	workTree = initFlags.String("work.tree", "~/", "The root location under which all files that you would like to save in git should be stored")
	sshKey = initFlags.String("ssh.key", "~/.ssh/id_rsa", "The SSH key used to interact with your private git repository storing your system config")

	if len(os.Args) == 0 {
		initFlags.PrintDefaults()
		os.Exit(0)
	}

	debugFlagSet := flag.NewFlagSet("debug", flag.ContinueOnError)
	debug = debugFlagSet.Bool("debug", false, "Output all errors as they are returned from the CLI to stdout")
	debugFlagSet.Parse(os.Args)

	switch os.Args[1] {
	case "help":
		fmt.Println(helpMessage)
		initFlags.PrintDefaults()
		debugFlagSet.PrintDefaults()
		os.Exit(0)
	case "version", "v":
		fmt.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		os.Exit(0)
	case "display":
		fmt.Println(fmt.Sprint("Config version: ", Version, ", Build time: ", BuildTime, ", Git hash: ", GitHash))
		system := config.System{}
		cli := config.NewCLI(*debug, log, system)
		conf := cli.LoadGitConfig()

		contents, _ := yaml.Marshal(&conf)
		fmt.Println(string(contents))
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
		system := config.System{}
		cli := config.NewCLI(*debug, log, system)

		file := fmt.Sprint(config.ConfigDir, "/", config.GitConfig)
		cli.UpdateConfigFile(file, conf)

		scm := config.NewGit(*debug, log, system)
		scm.Init(*repository, conf.GitDir)
		scm.ConfigureSSHKey(conf.GitDir, *sshKey)
		scm.IgnoreUnknownFiles(conf.GitDir, true)
	default:
		system := config.System{}
		cli := config.NewCLI(*debug, log, system)
		conf := cli.LoadGitConfig()

		if conf.GitDir == "" || conf.WorkTree == "" {
			fmt.Println("To initialise the CLI, run `config init` with the flags shown below:")
			initFlags.PrintDefaults()
			debugFlagSet.PrintDefaults()
			os.Exit(1)
		}

		params := make([]string, 0)
		params = append(params, "--git-dir", conf.GitDir)
		params = append(params, "--work-tree", conf.WorkTree)
		params = append(params, os.Args[1:]...)
		system.Run(
			"git",
			params...,
		)
	}
}
