
# history configuration
HISTFILE=$ZDOTDIR/.zsh_histfile
HISTSIZE=1000
SAVEHIST=100000
setopt appendhistory histignorespace
setopt inc_append_history hist_ignore_dups

# function to parse the history file and set the prompt
function retrieve-history() {
    local _prompt="$(thelm -t history -fP < "$HISTFILE")"
    zle reset-prompt
    BUFFER="${_prompt#*;}"
    zle end-of-line
}
zle -N retrieve-history

# Set the key binding to Meta-,
bindkey "^[," retrieve-history
