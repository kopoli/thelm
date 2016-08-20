package thelm

import (
	"bytes"
	"fmt"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

type ui struct {
	view  UIView

	cmd Command

	input  string
	cursor int
}

// UiAbortedErr tells if user wanted to abort
var UiAbortedErr = E.New("User interface was aborted")

var UiSelectedErr = E.New("User selected a line")

// Input line handling

// adds character at the cursor position
func (u *ui) addInputRune(ch rune) {
	var buf bytes.Buffer

	buf.WriteString(u.input[:u.cursor])
	buf.WriteRune(ch)
	buf.WriteString(u.input[u.cursor:])
	u.cursor++
	u.input = buf.String()
}

// removes count characters from cursor
func (u *ui) removeInput(count int) {
	var buf bytes.Buffer

	start := minmax(0, u.cursor-count, len(u.input))
	buf.WriteString(u.input[:start])
	buf.WriteString(u.input[u.cursor:])
	u.input = buf.String()
	u.cursor = start
}

// updates the statusline
func (u *ui) setStatusLine(lines int) {
	u.view.SetStatusLine(fmt.Sprintf(" Thelm - %d", lines))
}

// EditInput handles the input line manipulation
func (u *ui) EditInput(ev termbox.Event) error {

	if ev.Ch != 0 {
		u.addInputRune(ev.Ch)
	} else {
		switch ev.Key {
		case termbox.KeyArrowLeft:
			u.cursor--
		case termbox.KeyArrowRight:
			u.cursor++
		case termbox.KeySpace:
			u.addInputRune(' ')
		case termbox.KeyBackspace:
			u.removeInput(1)
		case termbox.KeyBackspace2:
			u.removeInput(1)
		default:
			return nil
		}
	}

	// Update the input line
	u.cursor = minmax(0, u.cursor, len(u.input))
	u.view.SetInputLine(u.input, u.cursor)
	u.view.Flush()

	// Run the command
	u.view.Clear()
	args := strings.Split(u.input, " ")
	err := u.cmd.Run(args[0], args[1:]...)
	if err != nil {
		u.setStatusLine(0)
	}
	return nil
}

// View / General key handling

// Keybinding handler
type handlerFunc func(termbox.Key) error

func (u *ui) moveCursor(ydiff int) error {
	u.view.MoveHighlightLine(ydiff)
	u.view.Flush()
	return nil
}

func (u *ui) cmdSelectUp(termbox.Key) error {
	return u.moveCursor(-1)
}

func (u *ui) cmdSelectDown(termbox.Key) error {
	return u.moveCursor(1)
}

// func (u *ui) cmdSelectPgUp(termbox.Key) error {
// 	return u.moveCursorPage(g, -1)
// }

// func (u *ui) cmdSelectPgDown(termbox.Key) error {
// 	return u.moveCursorPage(g, 1)
// }

// func (u *ui) cmdToggleDebug(g *gocui.Gui, v *gocui.View) (err error) {
// }

// func (u *ui) cmdSelectLine(g *gocui.Gui, v *gocui.View) (err error) {

// }

func (u *ui) cmdSelectLine(termbox.Key) error {
	return UiSelectedErr
}

func (u *ui) cmdAbort(termbox.Key) error {
	return UiAbortedErr
}

func (u *ui) handleEventKey(key termbox.Key) (err error) {

	keyHandlers := map[termbox.Key]handlerFunc{
		termbox.KeyCtrlG:     u.cmdAbort,
		termbox.KeyArrowUp:   u.cmdSelectUp,
		termbox.KeyArrowDown: u.cmdSelectDown,
		termbox.KeyEnter:     u.cmdSelectLine,
	}

	if handler, ok := keyHandlers[key]; ok {
		err = handler(key)
	}
	return
}

// The handling of the command running

// Write receives data to be displayed on screen
func (u *ui) Write(p []byte) (n int, err error) {
	n, err = u.view.Write(p)
	u.setStatusLine(u.view.GetDataLines())
	u.view.Flush()
	return
}

// Ui runs the user interface that selects the line from input
func Ui(opts Options, args []string) (ret string, err error) {

	var u ui

	// Termbox setup
	err = termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// Testing ui.
	u.view.SetStatusLine(" Thelm testi ")
	fmt.Fprintln(&u.view, "viewtesti")
	fmt.Fprintln(&u.view, "Tahan toiselle riville")
	fmt.Fprintln(&u.view, "Kolmas rivi")
	u.view.Flush()

	u.cmd.Passthrough = &u
	u.cmd.Run("ls", "-lah")

	// Main loop
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			default:
				err = u.handleEventKey(ev.Key)
				if err != nil {
					if err == UiSelectedErr {
						err = nil
						ret = u.view.GetHighlightLine()
					}
					return
				}

				err = u.EditInput(ev)
				if err != nil {
					return
				}
			}
		case termbox.EventError:
			err = ev.Err
			return
		}
	}

	return
}
