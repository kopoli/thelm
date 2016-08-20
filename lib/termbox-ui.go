package thelm

import (
	"bytes"
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

type ui struct {
	view UIView

	input  string
	cursor int
}

// UiAbortedErr tells if user wanted to abort
var UiAbortedErr = E.New("User interface was aborted")

var UiSelectedErr = E.New("User selected a line")

// Input line handling

// add character at the cursor position
func (u *ui) addInputRune(ch rune) {
	var buf bytes.Buffer

	buf.WriteString(u.input[:u.cursor])
	buf.WriteRune(ch)
	buf.WriteString(u.input[u.cursor:])
	u.cursor++
	u.input = buf.String()
}

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
		}
	}

	u.cursor = minmax(0, u.cursor, len(u.input))

	u.view.SetInputLine(u.input, u.cursor)
	u.view.Flush()
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

	// Testingui.
	u.view.SetStatusLine(" Thelm testi ")
	fmt.Fprintln(&u.view, "viewtesti")
	fmt.Fprintln(&u.view, "Tahan toiselle riville")
	fmt.Fprintln(&u.view, "Kolmas rivi")
	u.view.Flush()

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
