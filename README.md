# Thelm

Selection narrowing tool for the terminal. The basic idea is the adapted
from [Emacs-helm](https://github.com/emacs-helm/helm).

To see thelm in action, the following use cases contain animations:
- [Running ag or git-grep continuously](#running-ag-or-git-grep-continuously)
- [Insert a command to prompt from history](#insert-a-command-to-prompt-from-history)
- [Choose directory where to jump from directory stack](#choose-directory-where-to-jump-from-directory-stack)

## Installation

```
$ go get github.com/kopoli/thelm
```

Alternatively you can download one of the pre-compiled releases.


## Usage

Basic operation is the following:

- Reads data from:
  - Running given command repeatedly.
  - A pipe.
  - A file.

- User is able to select a line visually with a
  [termbox-go](https://github.com/nsf/termbox-go) user interface.

- The line is printed out to stdout.

### Command line

```
$ ./thelm --help

Usage: ./thelm [OPTIONS] [-- ARG...]

Helm for terminal

Arguments:
  ARG                         Command to be run

Options:
  -v, --version               Show the version and exit
  -f, --filter                Start filtering after running command.
  -d, --default               The default argument that will be printed out if aborted.
  -i, --hide-initial          Hide command given at the command line.
  -s, --single-arg            Regard input given in the UI as a single argument to the program.
  -r, --relaxed-regexp        Regard input as a relaxed regexp. Implies --single-arg.
  -t, --title                 Title string in UI. (default "./thelm")
  -F, --file                  The file which will be read instead of running a command.
  -P, --pipe                  The data will be read through a pipe.
      --licenses              Show licenses of ./thelm.
      --cpu-profile-file      The CPU profile would be saved to this file.
      --memory-profile-file   The Memory profile would be saved to this file.
```

Concepts:
- **Filtering** means that the current input can be filtered/narrowed down
  grep -style. This can be toggled with a key binding. If the `-f` flag is
  given, the data is read once (or a command is run once) and then it is
  filtered within thelm.

- **Default argument** is what is printed out if user aborts out of the UI. If
  this is not supplied, nothing is printed.

- **Hiding initial arguments** means that the user can give a part of the
  command from the command line and the rest in the UI. See for an example
  in
  [Running ag or git-grep continuously](#running-ag-or-git-grep-continuously).

- **Single argument** is used when the part of the command line given in the
  UI is wanted to be given to the called program as a single argument. Without
  this, whitespace delimited fields are interpret as separate arguments.

- **Reading from a file** is invoked with the `-F` flag.

- **Reading from a pipe** is invoked with the `-P` flag.

### Key bindings within the narrowing UI

| Key      | Function                                           |
|----------|----------------------------------------------------|
| ESC      | Abort.                                             |
| C-g      | Abort.                                             |
| up C-p   | Select previous item.                              |
| down C-n | Select next item.                                  |
| pgup     | Move focus backward a screenful.                   |
| pgdown   | Move focus forward a screenful.                    |
| RET      | Terminate UI and print highlighted line to stdout. |
| C-f      | Toggle filtering.                                  |


## Use cases

### Insert a command to prompt from history

This is comparable to the "C-r" binding in bash and zsh. I have only zsh
configuration for this setup. It can be found
[here](use-cases/retrieve-history/zsh). This is configured to "Meta-,"
keybinding by default.

![insert command animation](https://github.com/kopoli/thelm/raw/master/use-cases/retrieve-history/animation.gif)

### Choose directory where to jump from directory stack

The ZSH configuration is in [here](use-cases/jump-dirstack/zsh). It uses
"Meta-." key binding by default.

![jump dirstack animation](https://github.com/kopoli/thelm/raw/master/use-cases/jump-dirstack/animation.gif)

### Running ag or git-grep continuously

```
$ ./thelm --hide-initial --single-arg ag
```

![continuous ag animation](https://github.com/kopoli/thelm/raw/master/use-cases/continuous-ag/animation.gif)

This is marginally useful from the terminal. Obviously I prefer
[helm-grepint](https://github.com/kopoli/helm-grepint) for Emacs.


## License

MIT license
