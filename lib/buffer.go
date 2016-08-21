package thelm

import (
	"bytes"
	"io"
	"regexp"
)

type Buffer struct {
	// lines [][]byte
	buffer []byte
	readpos int

	// Callback that provides data out
	Passthrough io.Writer
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.buffer = append(b.buffer, p...)
	n = len(p)
	return
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	n = copy(p, b.buffer[b.readpos:])
	b.readpos += n
	if b.readpos == len(b.buffer) {
		err = io.EOF
	}
	return
}

// Filter the current Buffer with the regexp and write output to out.
func (b *Buffer) Filter(regex string) (err error) {
	var lines int
	re, err := regexp.Compile("(?i)" + AsRelaxedRegexp(regex))
	if err != nil {
		return
	}

	for _, line := range bytes.Split(b.buffer, []byte("\n")) {
		if re.Match(line) {
			lines += 1
			_, err = b.Passthrough.Write(append(line, '\n'))
			if err != nil {
				return
			}
		}
	}

	if lines == 0 {
		err = E.New("No lines matching the filter found")
	}
	return
}
