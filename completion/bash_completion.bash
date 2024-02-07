#/usr/bin/env bash
#
# COMP_WORDS: an array of all the words typed after the name of the program the compspec belongs to;
# COMP_CWORD: the index of the word the cursor was when the tab key was pressed;
# COMP_LINE: the current command line.

# complete -W "init completion git help" config

_config_completions()
{
  # COMPREPLY+=("init")
  # COMPREPLY+=("completion")
  # COMPREPLY+=("git")
  # COMPREPLY+=("help")
	COMPREPLY=($(compgen -W "init completion git help" "${COMP_WORDS[1]}"))
}

complete -F _config_completions config
