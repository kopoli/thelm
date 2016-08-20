package thelm

import (
	"bytes"
	"sync"

	"github.com/nsf/termbox-go"
)

// UIView is the graphical representation of the data. It is used to draw the
// data to the screen.
type UIView struct {
	buffer           []byte
	lines            int
	sizeX, sizeY     int
	offsetX, offsetY int
	highlightY       int

	statusLine string

	inputLine    string
	inputCursorX int

	mutex sync.Mutex
}

// gets the offset of the start of the next line
func (u *UIView) nextLineOffset(start int) (offset int) {
	offset = bytes.Index(u.buffer[start:], []byte("\n")) + 1

	if offset == 0 {
		offset = len(u.buffer)
	} else {
		offset += start
	}
	return offset
}

// gets the byte offset of a given line
func (u *UIView) lineToByteOffset(inputLine int) (offset int) {
	for line := 0; line < inputLine; line++ {
		offset = u.nextLineOffset(offset)
	}
	return
}

// gets the offset of the first line that should be drawn
func (u *UIView) startLineOffset() (offset int) {
	return u.lineToByteOffset(u.offsetY)
}

// draws a string on screen
func (u *UIView) drawText(x, y int, fg, bg termbox.Attribute, text string) {
	for _, ch := range text {
		termbox.SetCell(x, y, ch, fg, bg)
		x++
	}
}

// fills a line on screen with given cell
func (u *UIView) fillLine(x, y, w int, fg, bg termbox.Attribute, ch rune) {
	for pos := 0; pos < w; pos++ {
		termbox.SetCell(x+pos, y, ch, fg, bg)
	}
}

// update the view size which is the screen minus the input and the data lines
func (u *UIView) updateViewSize() {
	u.sizeX, u.sizeY = termbox.Size()
	u.sizeY -= 2
}

func minmax(low int, value int, high int) int {
	if value < low {
		return low
	} else if value > high {
		return high
	}
	return value
}

// The public interface

// Write writes the given data to the view. This can be called from anywhere.
func (u *UIView) Write(p []byte) (n int, err error) {
	u.mutex.Lock()

	u.lines += bytes.Count(p, []byte("\n"))
	u.buffer = append(u.buffer, p...)
	n = len(p)

	u.mutex.Unlock()
	return
}

// Clear clears the view buffer
func (u *UIView) Clear() {
	u.mutex.Lock()
	u.buffer = []byte{}
	u.lines = 0
	u.highlightY = 0
	u.offsetX = 0
	u.offsetY = 0
	u.mutex.Unlock()
	u.ShiftView(0, 0)
}

// Flush updates the whole screen
func (u *UIView) Flush() {
	u.mutex.Lock()

	// Clear the screen
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	u.updateViewSize()

	// Get the start line in the buffer
	if u.lines > 0 {
		pos := u.offsetX
		if u.offsetY > 0 {
			pos += u.startLineOffset()
		}

		fg := coldef
		bg := coldef

		// Draw the buffer text on screen
		for y := 0; y < u.sizeY; y++ {
			end := u.nextLineOffset(pos) - 1

			// Set up highlighting
			if u.highlightY == y {
				bg = termbox.AttrReverse
				fg = termbox.AttrBold
			} else {
				fg = coldef
				bg = coldef
			}

			// fmt.Println("pos", pos, "end", end, "len", len(u.buffer))
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
	}

	// Draw the statusline
	y := u.sizeY - 1
	u.fillLine(0, y, 2, coldef, coldef, '-')
	u.drawText(2, y, coldef, coldef, u.statusLine)
	pos := 2 + len(u.statusLine)
	u.fillLine(pos, y, u.sizeX-pos, coldef, coldef, '-')

	// Draw the input line
	y = u.sizeY
	u.drawText(0, y, coldef, coldef, u.inputLine)
	termbox.SetCursor(u.inputCursorX, y)

	termbox.Flush()
	u.mutex.Unlock()
}

// SetStatusLine sets the status line to a given string
func (u *UIView) SetStatusLine(line string) {
	u.mutex.Lock()
	u.statusLine = line
	u.mutex.Unlock()
}

// MoveInputLine sets the input line to given and moves the cursor to absolute x
func (u *UIView) SetInputLine(line string, cursorx int) {
	u.mutex.Lock()
	u.inputLine = line
	u.inputCursorX = minmax(0, cursorx, len(line))
	u.mutex.Unlock()
}

// HighlightLine highlights a given line on the view.
func (u *UIView) MoveHighlightLine(ydiff int) {
	u.mutex.Lock()
	lines := u.lines - 1
	if lines < 0 {
		lines = 0
	}
	u.highlightY = minmax(0, u.highlightY+ydiff, lines)
	u.mutex.Unlock()
}

// GetHighlightLine returns the string from the view buffer that is currently
// highlighted
func (u *UIView) GetHighlightLine() string {
	u.mutex.Lock()

	start := u.lineToByteOffset(u.highlightY)
	stop := u.nextLineOffset(start)

	ret := string(u.buffer[start:stop])

	u.mutex.Unlock()
	return ret
}

// ShiftView shifts the view by given difference in the data
func (u *UIView) ShiftView(x int, y int) {
	u.mutex.Lock()
	u.offsetX = x
	u.offsetY = y
	u.mutex.Unlock()
}

// ViewSize returns the size of the view
func (u *UIView) ViewSize() (x int, y int) {
	u.mutex.Lock()
	u.updateViewSize()
	u.mutex.Unlock()
	return u.sizeX, u.sizeY
}

// GetDataLines returns the number of lines in the view
func (u *UIView) GetDataLines() (ret int) {
	u.mutex.Lock()
	ret = u.lines
	u.mutex.Unlock()
	return
}
