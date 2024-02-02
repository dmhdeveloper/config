#compdef cfg

typeset -A opt_args

_arguments -C \
	'1:cmd:->cmds' \
	'*:: :->args' &&
	ret=0

case "$state" in
cmds)
	local commands
	commands=(
		'init:Initialise and configure cfg cli'
		'git:Invoke the git cli to manage configuration'
		'completion:Install completion scripts'
	)

	_describe -t commands 'command' commands && ret=0
	;;
args)
	case $line[1] in
	init)
		_arguments \
			'--url[The git remote URL (SSH) for your configuration files]::<URL>' \
			'--git.dir[The directory to store your .git files (bare repo)]:directory:_files -/ -g ~/.dotfiles' \
			'--work.dir[The root directory in which all configuration files can be found]:directory:_files -/ -g ~/' \
			'--ssh.key[The SSH key used to access the git remote repo storing your configuration]:directory:_files -g ~/.ssh/id_rsa' && \
			ret=0
		;;
	completion)
		local commands
		commands=(
			'zsh:Install zsh completion script'
		)

		_describe -t commands 'command' commands && ret=0
		;;
	git)
		local service=git
		_git
		;;
	esac
	;;
esac

return 1
