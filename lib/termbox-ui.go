package thelm

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kopoli/appkit"
	termbox "github.com/nsf/termbox-go"
)

type filter struct {
	buf         Buffer
	savedInput  string
	savedCursor int
}

type ui struct {
	optInputTitle string
	optSingleArg  bool
	optRelaxedRe  bool

	hiddenArgs []string

	view UIView

	// cmd Command
	source AsyncSource

	input  string
	cursor int

	filter *filter
}

// UiAbortedErr tells if user wanted to abort
var UiAbortedErr = E.New("User interface was aborted")

// UiSelectedErr tells that user selected a line
var UiSelectedErr = E.New("User selected a line")

// Input line handling

// adds character at the cursor position
func (u *ui) addInputRune(ch rune) {
	var buf bytes.Buffer

	input := []rune(u.input)

	buf.WriteString(string(input[:u.cursor]))
	buf.WriteRune(ch)
	buf.WriteString(string(input[u.cursor:]))
	u.cursor++
	u.input = buf.String()
}

// removes count characters from cursor
func (u *ui) removeInput(count int) {
	var buf bytes.Buffer

	input := []rune(u.input)

	start := minmax(0, u.cursor-count, len(input))
	buf.WriteString(string(input[:start]))
	buf.WriteString(string(input[u.cursor:]))
	u.input = buf.String()
	u.cursor = start
}

// clears the whole input string
func (u *ui) clearInput() {
	u.input = ""
	u.cursor = 0
}

// removes characters up to thep previous space
func (u *ui) backwardKillWord() {
	lastspace := strings.LastIndex(u.input[:u.cursor], " ")
	if lastspace < 0 {
		lastspace = 0
	}

	u.removeInput(u.cursor - lastspace)
}

// updates the statusline
func (u *ui) setStatusLine(lines int) {
	status := ""
	if u.filter != nil {
		status = " - filtering"
	}
	u.view.SetStatusLine(fmt.Sprintf(" %s%s - %d ", u.optInputTitle, status,
		lines))
}

// Runs the command that has been stored in input and hiddenArgs
func (u *ui) RunCommand() {
	if !u.source.IsOneShot() {
		// Finish the possibly previous command run
		u.source.Finish()

		u.view.Clear()
	}

	line := u.input
	var args []string

	args = append(args, u.hiddenArgs...)

	if u.optRelaxedRe {
		line = AsRelaxedRegexp(line)
	}

	if u.optSingleArg {
		args = append(args, line)
	} else {
		args = append(args, strings.Split(line, " ")...)
	}

	err := u.source.Run(args...)
	if err != nil {
		u.setStatusLine(0)
	}
}

// Refresh updates the UI from the internal data
func (u *ui) Refresh(update bool) {
	// Update the input line
	u.cursor = minmax(0, u.cursor, len(u.input))
	u.view.SetInputLine(u.input, u.cursor)
	u.view.Flush()

	if !update {
		return
	}

	u.setStatusLine(0)

	// Generate the output
	if u.filter != nil {
		u.view.Clear()
		u.filter.buf.Filter(u.input)
	} else {
		u.RunCommand()
	}

}

// EditInput handles the input line manipulation
func (u *ui) EditInput(ev termbox.Event) error {
	update := true

	// Visible character input
	if ev.Ch != 0 {
		u.addInputRune(ev.Ch)
	} else {
		key := ev.Key
		mod := ev.Mod

		// Keys
		switch {
		case key == termbox.KeyArrowLeft:
			u.cursor--
			update = false
		case key == termbox.KeyArrowRight:
			u.cursor++
			update = false
		case key == termbox.KeySpace:
			u.addInputRune(' ')
		case key == termbox.KeyBackspace || key == termbox.KeyBackspace2:
			u.removeInput(1)
		case key == termbox.KeyCtrlU:
			u.clearInput()
		case (mod == termbox.ModAlt && (key == termbox.KeyBackspace ||
			key == termbox.KeyBackspace2)) ||
			key == termbox.KeyCtrlW:
			u.backwardKillWord()
		default:
			return nil
		}
	}

	u.Refresh(update)

	return nil
}

// View / General key handling

// Keybinding handler
type handlerFunc func(termbox.Key) error

func (u *ui) moveCursor(ydiff int) error {
	u.view.MoveHighlightAndView(ydiff)
	u.view.Flush()
	return nil
}

func (u *ui) moveCursorPage(ydiff int) error {
	_, ypage := u.view.ViewSize()
	u.moveCursor(ydiff * ypage)
	return nil
}

func (u *ui) cmdSelectUp(termbox.Key) error {
	return u.moveCursor(-1)
}

func (u *ui) cmdSelectDown(termbox.Key) error {
	return u.moveCursor(1)
}

func (u *ui) cmdSelectPgUp(termbox.Key) error {
	return u.moveCursorPage(-1)
}

func (u *ui) cmdSelectPgDown(termbox.Key) error {
	return u.moveCursorPage(1)
}

func (u *ui) cmdSelectLine(termbox.Key) error {
	return UiSelectedErr
}

func (u *ui) cmdAbort(termbox.Key) error {
	return UiAbortedErr
}

func (u *ui) cmdToggleFilter(termbox.Key) error {

	if u.filter == nil {
		u.filter = &filter{
			savedInput:  u.input,
			savedCursor: u.cursor,
		}

		u.input = ""
		u.cursor = 0
		u.filter.buf.Passthrough = u
		io.Copy(&u.filter.buf, &u.view)
	} else {
		_, line := u.view.GetHighlightLine()
		u.view.Clear()
		io.Copy(&u.view, &u.filter.buf)
		u.view.MoveHighlightAndView(u.filter.buf.GetRealLine(line))
		u.input = u.filter.savedInput
		u.cursor = u.filter.savedCursor
		u.filter.buf.Close()
		u.filter = nil
	}
	u.setStatusLine(0)
	u.view.SetInputLine(u.input, u.cursor)
	u.view.Flush()

	return nil
}

func (u *ui) handleEventKey(key termbox.Key) (err error) {

	keyHandlers := map[termbox.Key]handlerFunc{
		termbox.KeyEsc:       u.cmdAbort,
		termbox.KeyCtrlG:     u.cmdAbort,
		termbox.KeyArrowUp:   u.cmdSelectUp,
		termbox.KeyArrowDown: u.cmdSelectDown,
		termbox.KeyCtrlP:     u.cmdSelectUp,
		termbox.KeyCtrlN:     u.cmdSelectDown,
		termbox.KeyPgdn:      u.cmdSelectPgDown,
		termbox.KeyPgup:      u.cmdSelectPgUp,
		termbox.KeyEnter:     u.cmdSelectLine,
		termbox.KeyCtrlF:     u.cmdToggleFilter,
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
	u.setStatusLine(u.view.GetDataLineCount())
	u.view.Flush()
	return
}

// Ui runs the user interface that selects the line from input
func Ui(opts appkit.Options, args []string) (ret string, err error) {

	var u ui

	u.optInputTitle = opts.Get("input-title", "thelm")
	u.optSingleArg = opts.IsSet("single-argument")
	u.optRelaxedRe = opts.IsSet("relaxed-regexp")

	if opts.IsSet("hide-initial-args") {
		u.hiddenArgs = args
	} else {
		u.input = strings.Join(args, " ")
		u.cursor = len(u.input)
	}

	// Set input source
	enableFiltering := opts.IsSet("enable-filtering")
	inputPipe := opts.IsSet("input-pipe")
	inputFile := opts.Get("input-file", "")

	if inputPipe || inputFile != "" {
		if inputPipe {
			u.source = &SourceReader{
				Input: os.Stdin,
			}
		} else {
			u.source = &SourceFile{
				FileName: inputFile,
			}
		}
		enableFiltering = true
		u.input = ""
	} else {
		// Command setup
		u.source = &Command{}
	}

	u.source.SetOutput(&u)

	// Termbox setup
	err = termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// Set up the ui and initial draw
	u.setStatusLine(0)
	u.Refresh(true)

	if enableFiltering {
		u.source.Wait()
		u.cmdToggleFilter(termbox.KeyCtrlF)
		u.Refresh(true)
	}

	// Main loop
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			err = u.handleEventKey(ev.Key)
			if err != nil {
				if err == UiSelectedErr {
					err = nil
					ret, _ = u.view.GetHighlightLine()
				}
				return
			}

			err = u.EditInput(ev)
			if err != nil {
				return
			}
		case termbox.EventError:
			err = ev.Err
			return
		}
	}

	return
}
