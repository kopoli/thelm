package thelm

import (
	"bytes"
	"io"
	"regexp"
)

type Buffer struct {
	data []byte

	// Callback that provides data out
	Sync func([]byte) error
}

func (b *Buffer) Push(data string) {
	b.data = []byte(data)
}

func (b *Buffer) Pop(out io.Writer) (err error) {
	_, err = out.Write(b.data)
	return
}

// Filter the current Buffer with the regexp and write output to out.
func (b *Buffer) Filter(regex string) (lines int, err error) {
	re, err := regexp.Compile("(?i)" + AsRelaxedRegexp(regex))
	if err != nil {
		return
	}

	//TODO Splitting can be done in the Push phase
	for _, line := range bytes.Split(b.data, []byte("\n")) {
		if re.Match(line) {
			lines += 1
			// _, err = out.Write(append(line, '\n'))
			err = b.Sync(append(line, '\n'))
			if err != nil {
				return
			}
		}
	}

	return
}
