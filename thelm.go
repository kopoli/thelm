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
	opts.Set("program-name", "thelm")
	opts.Set("program-version", progVersion)

	line, err := thelm.Ui(opts)

	if err == thelm.UiAbortedErr {
		os.Exit(1)
	}
	if err != nil {
		fault(err, "Running user interface failed")
	}

	fmt.Print(line)
	os.Exit(0)
}
