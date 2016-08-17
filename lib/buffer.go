package thelm

import (
	"bytes"
	"io"
	"regexp"
)

type Buffer struct {
	lines [][]byte

	// Callback that provides data out
	Sync func([]byte) error
}

func (b *Buffer) Push(data string) {
	b.lines = bytes.Split([]byte(data), []byte("\n"))
}

func (b *Buffer) Pop(out io.Writer) (err error) {
	data := bytes.Join(b.lines, []byte("\n"))
	_, err = out.Write(data)
	return
}

// Filter the current Buffer with the regexp and write output to out.
func (b *Buffer) Filter(regex string) (lines int, err error) {
	re, err := regexp.Compile("(?i)" + AsRelaxedRegexp(regex))
	if err != nil {
		return
	}

	for _, line := range b.lines {
		if re.Match(line) {
			lines += 1
			err = b.Sync(append(line, '\n'))
			if err != nil {
				return
			}
		}
	}

	return
}
