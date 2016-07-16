package thelm

import (
	"fmt"
	"runtime"

	"github.com/jawher/mow.cli"
)

func Cli(opts Options, argsin []string) (args []string, err error) {
	progName := opts.Get("program-name", "thelm")
	progVersion := opts.Get("program-version", "undefined")
	app := cli.App(progName, "Helm for terminal")

	app.Spec = "[OPTIONS] [-- ARG...]"

	app.Version("version v", fmt.Sprintf("%s: %s\nBuilt with: %s/%s on %s/%s",
		progName, progVersion, runtime.Compiler, runtime.Version(),
		runtime.GOOS, runtime.GOARCH))

	optFilter := app.BoolOpt("filter f", false, "Start after running command")

	argArg := app.StringsArg("ARG", nil, "Command to be run")
	app.Action = func() {
		args = *argArg
		if *optFilter {
			opts.Set("enable-filtering", "t")
		}
	}

	err = app.Run(argsin)
	return
}
