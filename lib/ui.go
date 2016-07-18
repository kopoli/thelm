package thelm

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

// UiAbortedErr tells if user wanted to abort
var UiAbortedErr = E.New("User interface was aborted")

type ui struct {
	args            []string
	hideInitialArgs bool
	singleArg       bool
	relaxedRe       bool

	cmd  Command
	gui  *gocui.Gui
	line string

	inputTitle string

	filter   *Buffer
	prevline string
}

func (u *ui) abort(g *gocui.Gui, v *gocui.View) error {
	return UiAbortedErr
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

	if y+oy > u.cmd.Out.Count {
		y = u.cmd.Out.Count - oy
	}

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
	if err != nil {
		return
	}
	return gocui.ErrQuit
}

func (u *ui) pushFilter(g *gocui.Gui, v *gocui.View) (err error) {
	if u.filter != nil {
		return
	}
	u.filter = &u.cmd.Out
	u.prevline, err = u.clearInput("")
	return
}

func (u *ui) popFilter(g *gocui.Gui, v *gocui.View) (err error) {
	if u.filter == nil {
		return
	}
	u.filter = nil
	u.clearInput(u.prevline)
	u.triggerRun()
	return
}

func (u *ui) keybindings() (err error) {
	binds := []struct {
		key interface{}
		f   func(*gocui.Gui, *gocui.View) error
	}{
		{gocui.KeyCtrlG, u.abort},
		{gocui.KeyCtrlC, u.abort},
		{gocui.KeyArrowDown, u.selectDown},
		{gocui.KeyCtrlN, u.selectDown},
		{gocui.KeyArrowUp, u.selectUp},
		{gocui.KeyCtrlP, u.selectUp},
		{gocui.KeyEnter, u.selectLine},
		{gocui.KeyCtrlF, u.pushFilter},
		{gocui.KeyCtrlU, u.popFilter},
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

		u.cmd.Setup(func() {
			g.Execute(func(g *gocui.Gui) (err error) {
				inp, err := g.View("input")
				filtering := ""
				if u.filter != nil {
					filtering = " - filtering"
				}
				inp.Title = fmt.Sprintf("%s%s - %d", u.inputTitle,
					filtering, u.cmd.Out.Count)
				return
			})
		}, v)
	}
	if err != nil {
		return
	}

	v, err = g.SetView("input", -1, maxy-2, maxx, maxy)
	if err == gocui.ErrUnknownView {
		v.Editable = true
		v.Title = u.inputTitle
		v.Wrap = true
		err = nil

		if !u.hideInitialArgs {
			initial := strings.Join(u.args, " ")
			fmt.Fprint(v, initial)
			v.SetCursor(len(initial), 0)
		}
		g.SetCurrentView("input")
	}
	if err != nil {
		return
	}

	return
}

func (u *ui) clearInput(in string) (out string, err error) {
	v, err := u.getInput()
	if err != nil {
		return
	}
	out, err = u.getInputLine()
	if err != nil {
		return
	}

	v.Clear()
	fmt.Fprint(v, in)
	v.SetCursor(len(in), 0)

	return
}

func (u *ui) getInput() (ret *gocui.View, err error) {
	ret, err = u.gui.View("input")
	if err != nil {
		err = E.Annotate(err, "Getting input view failed")
		return
	}
	return
}

func (u *ui) getInputLine() (ret string, err error) {
	v, err := u.getInput()
	if err != nil {
		return
	}
	ret, err = v.Line(0)
	if err != nil {
		err = E.Annotate(err, "Getting first input line failed")
		return
	}
	return
}

func (u *ui) triggerRun() (err error) {
	output, err := u.gui.View("output")
	if err != nil {
		err = E.Annotate(err, "Getting output view failed")
		return
	}

	u.cmd.Finish()
	output.Clear()
	output.SetCursor(0, 0)

	// Ignore error if input line cannot be read
	line, _ := u.getInputLine()

	if u.filter != nil {
		_ = u.filter.Filter(line)
	} else {
		args := make([]string, 0)
		if u.hideInitialArgs {
			args = append(args, u.args...)
		}

		if u.relaxedRe {
			line = AsRelaxedRegexp(line)
		}

		if u.singleArg {
			args = append(args, line)
		} else {
			args = append(args, strings.Split(line, " ")...)
		}

		err = u.cmd.Run(args[0], args[1:]...)
		if err != nil {
			err = E.Annotate(err, "Running command failed")
		}
	}
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

func Ui(opts Options, args []string) (ret string, err error) {

	var UI ui
	UI.inputTitle = opts.Get("input-title", "thelm")
	UI.hideInitialArgs = opts.IsSet("hide-initial-args")
	UI.singleArg = opts.IsSet("single-argument")
	UI.relaxedRe = opts.IsSet("relaxed-regexp")

	UI.args = args

	UI.gui = gocui.NewGui()
	err = UI.gui.Init()
	if err != nil {
		E.Annotate(err, "Initializing UI library failed")
		return
	}
	defer UI.gui.Close()
	defer UI.cmd.Finish()

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

	UI.gui.Execute(func(g *gocui.Gui) (err error) {
		err = UI.triggerRun()
		if err != nil {
			err = E.Annotate(err, "Initial run failed")
		}
		if opts.IsSet("enable-filtering") {
			UI.pushFilter(nil, nil)
		}
		return
	})

	err = UI.gui.MainLoop()
	ret = UI.line
	if err == gocui.ErrQuit {
		err = nil
	}
	if err == UiAbortedErr {
		return
	}
	if err != nil {
		E.Annotate(err, "Running UI main loop failed")
	}

	return
}
