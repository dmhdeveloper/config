# Config

Config is a small [Git][1] wrapper that uses a [Git Bare][2] repository for managing your local config files.
By setting some needed defaults and using a config file, `config` can monitor files from any root directory.
Once you have initialised the CLI, you interact with it the same way you would use Git, except the binary is `config`.

## Getting started

To install, either build from source by cloning the repo or using `go install`.

### Installing from source

```sh
git clone https://github.com/dmhdeveloper/config.git config
cd config
make build
cp config /usr/local/bin/
```

### Installing with go

```sh
go install github.com/dmhdeveloper/config/cmd/config@latest
```

### Configuring CLI

Config requires a URL which points to a git remote repository.
Interaction with your config repository is only allowed using a pre-configured SSH key.
The remote repository is meant to store your system config files so that you can manage their versions and persist changes like you would usually do with projects managed by Git.
Additional config is needed such as a `git-dir`, which is where the Git Bare repository files are kept, and a `work-tree`, which is the root directory that git is allowed to view.
To configure this, run the following commands.

```sh
config init -git.url "<remote-ssh-url>" -work.tree "<work-tree>" -git.dir "<bare-repo-dir>" -ssh.key "<ssh-key>"
```

## Roadmap

- [x] Install and configure CLI for interacting with git remote repository
- [ ] Support shell completion files
- [ ] Support Windows
- [ ] Support installing apps using package manager

[1]: https://git-scm.com/
[2]: https://git-scm.com/docs/git-init#Documentation/git-init.txt---bare
