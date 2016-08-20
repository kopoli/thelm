package thelm

import (
	"bytes"
	"sync"

	"github.com/nsf/termbox-go"
)

type UiView struct {
	buffer           []byte
	sizeX, sizeY     int
	offsetX, offsetY int
	highlightY       int

	statusLine string

	inputLine    string
	inputCursorX int

	mutex sync.Mutex
}

// gets the offset of the first line that should be drawn
func (u *UiView) startLineOffset() (offset int) {
	for line := 0; line < u.offsetY; line++ {
		offset = bytes.Index(u.buffer[offset:], []byte("\n")) + 1
	}
	return
}

// draws a string on screen
func (u *UiView) drawText(x, y int, fg, bg termbox.Attribute, text string) {
	for _, ch := range text {
		termbox.SetCell(x, y, ch, fg, bg)
		x++
	}
}

// fills a line on screen with given cell
func (u *UiView) fillLine(x, y, w int, fg, bg termbox.Attribute, ch rune) {
	for pos := 0; pos < w; pos++ {
		termbox.SetCell(x+pos, y, ch, fg, bg)
	}
}

// update the view size which is the screen minus the input and the data lines
func (u *UiView) updateViewSize() {
	u.sizeX, u.sizeY = termbox.Size()
	u.sizeY -= 2
}

// The public interface

// Write writes the given data to the view. This can be called from anywhere.
func (u *UiView) Write(p []byte) (n int, err error) {
	u.mutex.Lock()

	u.buffer = append(u.buffer, p...)
	n = len(p)

	u.mutex.Unlock()
	return
}

// Clear clears the view buffer
func (u *UiView) Clear() {
	u.mutex.Lock()
	u.buffer = []byte{}
	u.mutex.Unlock()
	u.ShiftView(0, 0)
}

// Flush updates the whole screen
func (u *UiView) Flush() {
	u.mutex.Lock()

	// Clear the screen
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	u.updateViewSize()

	// Get the start line in the buffer
	pos := u.offsetX
	if u.offsetY > 0 {
		pos += u.startLineOffset()
	}

	fg := coldef
	bg := coldef

	// Draw the buffer text on screen
	for y := 0; y < u.sizeY; y++ {
		end := bytes.Index(u.buffer[pos:], []byte("\n"))
		if end < 0 {
			end = len(u.buffer)
		} else {
			end += pos
		}

		// Set up highlighting
		if u.highlightY == y {
			bg = termbox.AttrReverse
			fg = termbox.AttrBold
		} else {
			fg = coldef
			bg = coldef
		}

		// fmt.Println("Pos", pos, "end", end, "len", len(u.buffer))

		line := string(u.buffer[pos:end])
		u.drawText(0, y, fg, bg, line)

		if u.highlightY == y {
			length := u.sizeX - len(line)
			u.fillLine(len(line), y, length, fg, bg, ' ')
		}

		pos = u.offsetX + end + 1
		if pos >= len(u.buffer) {
			break
		}
	}

	// Draw the statusline
	y := u.sizeY - 1
	u.fillLine(0, y, 2, coldef, coldef, '-')
	u.drawText(2, y, coldef, coldef, u.statusLine)
	pos = 2 + len(u.statusLine)
	u.fillLine(pos, y, u.sizeX-pos, coldef, coldef, '-')

	// Draw the input line
	y = u.sizeY
	u.drawText(0, y, coldef, coldef, u.inputLine)
	termbox.SetCursor(u.inputCursorX, y)

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

// ShiftView shifts the view by given difference in the data
func (u *UiView) ShiftView(x int, y int) {
	u.mutex.Lock()
	u.offsetX = x
	u.offsetY = y
	u.mutex.Unlock()
}

// ViewSize returns the size of the view
func (u *UiView) ViewSize() (x int, y int) {
	u.mutex.Lock()
	u.updateViewSize()
	u.mutex.Unlock()
	return u.sizeX, u.sizeY
}
