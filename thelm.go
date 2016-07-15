package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kopoli/thelm/lib"
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
	line, err := thelm.Ui(opts)

	if err != nil {
		fault(err, "Running user interface failed")
	}

	fmt.Print(line)
	os.Exit(0)
}
