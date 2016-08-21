package thelm

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/jawher/mow.cli"
)

// hideHelp hides the -h or --help argument if it appears after the -- argument
func hideHelp(args []string) (hidden bool, ret []string) {
	restpos := -1

	hidden = false
	ret = args

	for i, item := range args {
		if item == "--" {
			restpos = i
			break
		}
	}

	if restpos == -1 {
		return
	}

	for i, item := range args[restpos:] {
		if item == "-h" || item == "--help" {
			hidden = true
			args[i+restpos] = "//" + item
		}
	}
	return
}

func unhideHelp(args []string) []string {
	for i, item := range args {
		if item == "//-h" || item == "//--help" {
			args[i] = strings.TrimLeft(item, "/")
		}
	}
	return args
}

func Cli(opts Options, argsin []string) (args []string, err error) {
	progName := opts.Get("program-name", "thelm")
	progVersion := opts.Get("program-version", "undefined")
	app := cli.App(progName, "Helm for terminal")

	hidden, argsin := hideHelp(argsin)

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

	optCpuProfile := app.StringOpt("cpu-profile-file", "", "The CPU profile would be saved to this file.")
	optMemProfile := app.StringOpt("memory-profile-file", "", "The Memory profile would be saved to this file.")

	argArg := app.StringsArg("ARG", nil, "Command to be run")
	app.Action = func() {
		args = *argArg
		if hidden {
			args = unhideHelp(args)
		}
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

		opts.Set("cpu-profile-file", *optCpuProfile)
		opts.Set("memory-profile-file", *optMemProfile)
	}

	err = app.Run(argsin)
	return
}
