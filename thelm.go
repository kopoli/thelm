package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
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


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func layout(g *gocui.Gui) (err error) {
	maxx, maxy := g.Size()

	v, err := g.SetView("output", -1, -1, maxx, maxy - 2)
	if err == gocui.ErrUnknownView {
		v.Highlight = true
		fmt.Println(v, "Output goes here")
		err = nil
	}
	if err != nil {
		return
	}

	v, err = g.SetView("input", -1, maxy - 2, maxx, maxy)
	if err == gocui.ErrUnknownView {
		v.Editable = true
		v.Title = "jepajee"
		v.Wrap = true
		err = nil
		g.SetCurrentView("input")
	}
	if err != nil {
		return
	}

	return
}

func main() {

	gui := gocui.NewGui()
	err := gui.Init()
	if err != nil {
		fault(err, "Initializing UI library failed")
	}
	defer gui.Close()

	gui.SetLayout(layout)
	err = keybindings(gui)
	if err != nil {
		fault(err, "Setting keybindings failed")
	}

	gui.SelBgColor = gocui.ColorGreen
	gui.SelFgColor = gocui.ColorBlack
	gui.Cursor = true

	fmt.Println("TAALLA!!")

	err = gui.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		fault(err, "Running UI main loop failed")
	}
}
