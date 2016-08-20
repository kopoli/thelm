package thelm

import (
	"sync"

	"github.com/nsf/termbox-go"
)

type UiView struct {
	buffer           []byte
	offsetX, offsetY int
	highlightY       int

	statusLine string

	inputLine    string
	inputCursorX int

	mutex sync.Mutex
}

// The public interface

// Write writes the given data to the view. This can be called from anywhere.
func (u *UiView) Write(p []byte) (n int, err error) {
	u.mutex.Lock()

	n = len(p)
	u.mutex.Unlock()
	return
}

// Clear clears the view buffer
func (u *UiView) Clear() {
	u.mutex.Lock()

	u.mutex.Unlock()
}

// Flush updates the whole screen
func (u *UiView) Flush() {
	u.mutex.Lock()
	termbox.Flush()
	u.mutex.Unlock()
}

// SetStatusLine sets the status line to a given string
func (u *UiView) SetStatusLine(line string) {
	u.mutex.Lock()
	u.statusLine = line
	u.mutex.Unlock()
}

// SetInputLine sets the input line to given and moves the cursor to absolute x
func (u *UiView) SetInputLine(line string, cursorx int) {
	u.mutex.Lock()
	u.inputLine = line
	u.inputCursorX = cursorx
	u.mutex.Unlock()
}

// HighlightLine highlights a given line on the view
func (u *UiView) HighlightLine(y int) {
	u.mutex.Lock()
	u.highlightY = y
	u.mutex.Unlock()
}

// ShiftView shifts the view by given difference
func (u *UiView) ShiftView(x int, y int) {
	u.mutex.Lock()
	u.offsetX = x
	u.offsetY = y
	u.mutex.Unlock()
}

// ViewSize returns the size of the view
func (u *UiView) ViewSize() (x int, y int) {
	u.mutex.Lock()

	u.mutex.Unlock()
	return
}
