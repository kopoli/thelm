package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/kopoli/appkit"
	"github.com/kopoli/thelm/lib"
)

var (
	majorVersion     = "0"
	version          = "Undefined"
	timestamp        = "Undefined"
	progVersion      = majorVersion + "-" + version
	exitValue    int = 0
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

	// Exit goroutine and run all deferrals
	exitValue = 1
	runtime.Goexit()
}

func main() {
	opts := appkit.NewOptions()
	opts.Set("program-name", os.Args[0])
	opts.Set("program-version", progVersion)
	opts.Set("program-timestamp", timestamp)

	// In the last deferred function, exit the program with given code
	defer func() {
		os.Exit(exitValue)
	}()

	args, err := thelm.Cli(opts, os.Args)
	if err != nil {
		fault(err, "Parsing command line failed")
	}

	profiler, err := thelm.SetupProfiler(opts)
	if err != nil {
		fault(err, "Starting profiling failed")
	}
	defer profiler.Close()

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
		exitValue = 1
		return
	}
	if err != nil {
		fault(err, "Running user interface failed")
	}

	fmt.Println(line)
}
