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

	optFilter := app.BoolOpt("filter f", false, "Start filtering after running command.")
	optDefault := app.StringOpt("default d", "", "The default argument that will be printed out if aborted.")
	optHide := app.BoolOpt("hide-initial i", false, "Hide command given at the command line.")
	optSingleArg := app.BoolOpt("single-arg s", false, "Regard input given in the UI as a single argument to the program.")
	optRelaxedRe := app.BoolOpt("relaxed-regexp r", false, "Regard input as a relaxed regexp. Implies --single-arg.")
	optTitle := app.StringOpt("title t", progName, "Title string in UI.")

	argArg := app.StringsArg("ARG", nil, "Command to be run")
	app.Action = func() {
		args = *argArg
		if *optFilter {
			opts.Set("enable-filtering", "t")
		}
		if *optHide {
			opts.Set("hide-initial-args", "t")
		}
		if *optSingleArg {
			opts.Set("single-argument", "t")
		}
		if *optRelaxedRe {
			opts.Set("relaxed-regexp", "t")
		}

		opts.Set("default-value", *optDefault)
		opts.Set("input-title", *optTitle)
	}

	err = app.Run(argsin)
	return
}
