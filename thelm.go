package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kopoli/thelm/lib"
)

var (
	majorVersion = "0"
	version      = "Undefined"
	progVersion  = majorVersion + "-" + version
)

func printErr(err error, message string, arg ...string) {
	msg := ""
	if err != nil {
		msg = fmt.Sprintf(" (error: %s)", err)
	}
	fmt.Fprintf(os.Stderr, "Error: %s%s.%s\n", message, strings.Join(arg, " "), msg)
}

func fault(err error, message string, arg ...string) {
	printErr(err, message, arg...)
	os.Exit(1)
}

func main() {
	opts := thelm.GetOptions()
	opts.Set("program-name", os.Args[0])
	opts.Set("program-version", progVersion)
        retCode := 0

	args, err := thelm.Cli(opts, os.Args)
	if err != nil {
		fault(err, "Parsing command line failed")
	}

	var profiler thelm.Profiler

	err = profiler.Setup(opts)
	if err != nil {
		fault(err, "Could not set up profiling")
	}
	defer func() {
		profiler.Close()
		os.Exit(retCode)
	}()

	err = thelm.CheckSelfRunning(opts)
	if err != nil {
		fault(err, "Check that program isn't running in itself failed")
	}

	line, err := thelm.Ui(opts, args)
	if err == thelm.UiAbortedErr {
		defval := opts.Get("default-value", "")
		if defval != "" {
			fmt.Println(defval)
		}
		retCode = 1
		return
	}
	if err != nil {
		fault(err, "Running user interface failed")
	}

	fmt.Println(line)
}
