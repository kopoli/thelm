package thelm

import (
	"fmt"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
)

type ui struct {
	cmd  Command
	gui  *gocui.Gui
	line string
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func minmax(low int, value int, high int) int {
	if value < low {
		return low
	} else if value > high {
		return high
	}
	return value
}

func (u *ui) moveCursor(g *gocui.Gui, relpos int) (err error) {
	v, err := g.View("output")
	if err != nil {
		return
	}
	ox, oy := v.Origin()
	x, y := v.Cursor()
	_, maxy := v.Size()

	oy = oy + relpos
	y = y + relpos
	if oy > u.cmd.Out.Count {
		oy = u.cmd.Out.Count
	}
	if y < 0 || y >= maxy {
		if oy >= 0 {
			err = v.SetOrigin(ox, oy)
			if err != nil {
				return
			}
		}
		y = minmax(0, y, maxy-1)
	}

	// fmt.Println("cursor",y + oy, "count", tw.Count)
	// if y + oy > tw.Count {
	// 	y = oy - tw.Count
	// 	// y = 0
	// }

	err = v.SetCursor(x, y)
	return
}

func (u *ui) selectUp(g *gocui.Gui, v *gocui.View) error {
	return u.moveCursor(g, -1)
}

func (u *ui) selectDown(g *gocui.Gui, v *gocui.View) error {
	return u.moveCursor(g, 1)
}

func (u *ui) selectLine(g *gocui.Gui, v *gocui.View) (err error) {
	output, err := g.View("output")
	if err != nil {
		return
	}
	_, oy := output.Origin()
	_, y := output.Cursor()
	u.line, err = output.Line(y + oy)
	return gocui.ErrQuit
}

func (u *ui) keybindings() (err error) {

	binds := []struct {
		key interface{}
		f   func(*gocui.Gui, *gocui.View) error
	}{
		{gocui.KeyCtrlG, quit},
		{gocui.KeyCtrlC, quit},
		{gocui.KeyArrowDown, u.selectDown},
		{gocui.KeyCtrlN, u.selectDown},
		{gocui.KeyArrowUp, u.selectUp},
		{gocui.KeyCtrlP, u.selectUp},
		{gocui.KeyEnter, u.selectLine},
	}

	for _, b := range binds {
		err = u.gui.SetKeybinding("", b.key, gocui.ModNone, b.f)
		if err != nil {
			return
		}
	}

	return
}

func (u *ui) setLayout(g *gocui.Gui) (err error) {
	maxx, maxy := g.Size()

	v, err := g.SetView("output", -1, -1, maxx, maxy-2)
	if err == gocui.ErrUnknownView {
		v.Highlight = true
		err = nil

		u.cmd.Out = TriggeringWriter{
			Trigger: func() {
				g.Execute(func(g *gocui.Gui) (err error) {
					v, err := g.View("input")
					v.Title = fmt.Sprintf("thelm - %d", u.cmd.Out.Count)
					return
				})
			},
			Writer: v,
		}
	}
	if err != nil {
		return
	}

	v, err = g.SetView("input", -1, maxy-2, maxx, maxy)
	if err == gocui.ErrUnknownView {
		v.Editable = true
		v.Title = "thelm"
		v.Wrap = true
		err = nil
		fmt.Fprint(v, strings.Join(os.Args[1:], " "))
		g.SetCurrentView("input")
	}
	if err != nil {
		return
	}

	return
}

func (u *ui) getInputLine() (ret string, err error) {
	v, err := u.gui.View("input")
	if err != nil {
		E.Annotate(err, "Getting input view failed")
		return
	}
	ret, err = v.Line(0)
	if err != nil {
		E.Annotate(err, "Getting first input line failed")
		return
	}
	return
}

func (u *ui) triggerRun() (err error) {
	output, err := u.gui.View("output")
	if err != nil {
		E.Annotate(err, "Getting output view failed")
		return
	}
	output.Clear()
	output.SetCursor(0, 0)
	line, err := u.getInputLine()
	if err != nil {
		return
	}
	args := strings.Split(line, " ")
	fmt.Fprintln(output, "Command: ", line)
	err = u.cmd.Run(args[0], args[1:]...)
	return
}

// Edit implements the gocui.Editor interface. It is a single line editor.
func (u *ui) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case key == gocui.KeyArrowRight:
		x, _ := v.Cursor()
		line, _ := u.getInputLine()
		if x < len(line) {
			v.MoveCursor(1, 0, false)
		}
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
		u.triggerRun()
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
		u.triggerRun()
	case key == gocui.KeyDelete:
		v.EditDelete(false)
		u.triggerRun()

	// Disabled keys
	case key == gocui.KeyEnter:
	case key == gocui.KeyArrowDown:
	case key == gocui.KeyArrowUp:

	default:
		gocui.DefaultEditor.Edit(v, key, ch, mod)
	}
}

func Ui(opt Options) (ret string, err error) {

	var UI ui

	UI.gui = gocui.NewGui()
	err = UI.gui.Init()
	if err != nil {
		E.Annotate(err, "Initializing UI library failed")
		return
	}
	defer UI.gui.Close()

	UI.gui.Editor = &UI
	UI.gui.SelBgColor = gocui.ColorGreen
	UI.gui.SelFgColor = gocui.ColorBlack
	UI.gui.Cursor = true

	UI.gui.SetLayout(UI.setLayout)
	err = UI.keybindings()
	if err != nil {
		E.Annotate(err, "Setting keybindings failed")
		return
	}

	UI.gui.Execute(func(g *gocui.Gui) error {
		UI.triggerRun()
		return nil
	})

	err = UI.gui.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		E.Annotate(err, "Running UI main loop failed")
		return
	}

	err = nil
	ret = UI.line
	return
}
