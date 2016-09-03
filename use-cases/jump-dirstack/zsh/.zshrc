
# Directory handling
setopt autopushd pushd_ignore_dups
DIRSTACKSIZE=300
DIRSTACKFILE=$ZDOTDIR/.zsh-dirstack

function dirstack-load() {
    if [[ -f $DIRSTACKFILE ]]; then
        _dirstack=( ${(uf)"$(cat $DIRSTACKFILE)"} )
    fi
}

function dirstack-commit() {
    dirstack=(${(uf)_dirstack})
}

function dirstack-cleanup() {
    dirstack-load
    dirstack=()
    for d in $(echo $_dirstack); do
        test -d "$d" && { dirstack=( $dirstack "$d" ); }
    done
    dirstack-save
}

function dirstack-save() {
    dirs -pl >! $DIRSTACKFILE
}

# zsh callback which is run when changing directories
function chpwd() {
    dirstack-load
    dirstack=("${(@)_dirstack:#$PWD}") # remove pwd from dirstack.
    dirstack-save                      # pwd will be automatically the first in stack.
}


# function to jump to a directory
function jump-to-recent-dir() {
    local _tgtdir=$(thelm -t dirstack -fP < "$DIRSTACKFILE")
    if test -d "$_tgtdir"; then
        cd "$_tgtdir"
    fi
    zle reset-prompt
}
zle -N jump-to-recent-dir

# Set the key binding to Meta-.
bindkey "^[." jump-to-recent-dir

# Run these on start of the shell to initialize the directory stack
dirstack-load
dirstack-commit
