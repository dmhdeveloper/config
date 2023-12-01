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
go install github.com/dmhdeveloper/config@latest
```

## Roadmap

- [x] Implement storing and retrieving git config from config file
- [x] Implement command feature `init`
- [ ] Implement command feature `display`
- [ ] Implement command feature `app`
- [ ] Support shell completion files
- [ ] Support Windows

[1]: https://git-scm.com/
[2]: https://git-scm.com/docs/git-init#Documentation/git-init.txt---bare
