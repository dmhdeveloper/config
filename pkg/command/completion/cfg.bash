#!/bin/bash

# $1 The name of the command
# $2 The current word being completed (empty unless we are in the middle of typing a word).
# $3 The word before the word being completed.
# COMP_WORDS: an array of all the words typed after the name of the program the compspec belongs to;
# COMP_CWORD: the index of the word the cursor was when the tab key was pressed;
# COMP_LINE: the current command line.

# complete -W "init completion git help" config: adds words to completion
# complete -A <directory> config: adds directories to completion (can us a command to retrieve the current dirs or something similar)

# We need to find the last subcommand if there was one, ignoring the flags
# We will do this by working backwards through COMP_WORDS starting at COMP_CWORD until we reach a subcommand
# While going through the flags, add flags to a list of flags that have already been completed
# When we determine the subcommand, execute the subcommands function
# Once in the function, determine the current list of flags that have no been executed yet and return them
# Complete the flag if we only have one possibility
# When we have a flag that has not been populated, only completed, provide nothing for args that arent directories, or autocomplete dirs if needed.
# _cfg_subcommand()
# {
#
# }

_cfg_init() {
	case "${COMP_WORDS[$COMP_CWORD - 1]}" in
	--url) # Output the completion file for shell (bash|zsh)
		;;
	--git.dir) # Initialise the cfg cli
		COMPREPLY=($(compgen -W "~/.dotfiles" "${COMP_WORDS[$COMP_CWORD]}"))
		;;
	--work.tree) # Run git for managing config files
		COMPREPLY=($(compgen -W "~/" "${COMP_WORDS[$COMP_CWORD]}"))
		;;
	--ssh.key) # Add bash | zsh completion, at this point stop doing completion
		COMPREPLY=($(compgen -W "~/.ssh/id_rsa" "${COMP_WORDS[$COMP_CWORD]}"))
		;;
	*)
		flags=("--url" "--git.dir" "--work.tree" "--ssh.key")
		filtered=()
		for i in "${flags[@]}"; do
			skip=
			for j in "${COMP_WORDS[@]}"; do
				[[ $i == $j ]] && {
					skip=1
					break
				}
			done
			[[ -n $skip ]] || filtered+=("$i")
		done
		[[ ${#filtered[@]} -ne 0 ]] && COMPREPLY=($(compgen -W "${filtered[*]}" -- "${COMP_WORDS[COMP_CWORD]}"))
		;;
	esac
}

_cfg_completion() {
	COMPREPLY=($(compgen -W "bash zsh" "${COMP_WORDS[$COMP_CWORD]}"))
}

_cfg_git() {
	COMPREPLY=($(compgen -W "--help" "${COMP_WORDS[$COMP_CWORD]}"))
}

_cfg_completions() {
	if [ -z "$1" ]; then
		# display usage and exit when no args
		cfg -h
		return
	fi

	subcommand=""
	for i in "${COMP_WORDS[@]}"; do
		if [ "$i" == "$2" ]; then
			break
		fi
		if [ -z "$i" ]; then
			break
		fi
		if [[ "$i" == "-"* ]]; then
			break
		fi
		subcommand="$i"
	done

	shift
	case "$subcommand" in
	completion) # Output the completion file for shell (bash|zsh)
		_cfg_completion "$@"
		;;
	init) # Initialise the cfg cli
		_cfg_init "$@"
		;;
	git) # Run git for managing config files
		_cfg_git "$@"
		;;
	bash | zsh) # Add bash | zsh completion, at this point stop doing completion
		return $?
		;;
	*)
		subcommand="$1"
		COMPREPLY=($(compgen -W "init completion git help" "${COMP_WORDS[1]}"))
		;;
	esac
	return $?
}

complete -F _cfg_completions cfg
