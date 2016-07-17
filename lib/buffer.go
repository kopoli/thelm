package thelm

import (
	"bytes"
	"io"
	"regexp"
)

// Buffer stores data written to it through its io.Writer. It also writes the
// data to Passthrough. Each time Write is called it counts the lines to Count
// and runs the Trigger function
type Buffer struct {
	Passthrough io.Writer
	Count       int
	Trigger     func()
	data        []byte
}

func (b *Buffer) Reset() {
	b.Count = 0
	b.data = []byte{}
}

// Write data into the buffer and through to Passthrough.
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Count = b.Count + bytes.Count(p, []byte("\n"))
	b.data = append(b.data, p...)
	n, err = b.Passthrough.Write(p)
	b.Trigger()
	return
}

// Filter the current Buffer with the regexp and write output to Passthrough.
func (b *Buffer) Filter(regex string) (err error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return
	}

	for _, line := range bytes.Split(b.data, []byte("\n")) {
		if re.Match(line) {
			_, err = b.Passthrough.Write(append(line, '\n'))
			if err != nil {
				return
			}
		}
	}

	b.Trigger()

	return
}
