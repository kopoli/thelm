package thelm

import (
	"bytes"
	"io"
	"regexp"
)

type Buffer struct {
	buffer  []byte
	readpos int

	// Callback that provides data out
	Passthrough io.Writer

	done chan bool
}

// Write satisfies io.Writer. Writes data into the buffer.
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.buffer = append(b.buffer, p...)
	n = len(p)
	return
}

// Read satisfies io.Reader. Reads data from the buffer.
func (b *Buffer) Read(p []byte) (n int, err error) {
	n = copy(p, b.buffer[b.readpos:])
	b.readpos += n
	if b.readpos == len(b.buffer) {
		err = io.EOF
	}
	return
}

// Close  makes sure the asynchronous filtering is really stopped
func (b *Buffer) Close() (error) {
	if b.done != nil {
		b.done <- true
		close(b.done)
	}
	return nil
}

// Filter the current Buffer with the regexp and write output to out.
func (b *Buffer) Filter(regex string) (err error) {
	re, err := regexp.Compile("(?i)" + AsRelaxedRegexp(regex))
	if err != nil {
		return
	}

	if b.done == nil {
		b.done = make(chan bool)
	} else {
		b.done <- true
	}

	go func() {
		buf := []byte{}
		for _, line := range bytes.Split(b.buffer, []byte("\n")) {
			select {
			case <-b.done:
				return
			default:
			}
			if re.Match(line) {
				buf = append(buf, append(line, '\n')...)
			}
		}

		_, err = b.Passthrough.Write(buf)
		if err != nil {
			return
		}

		// Hang until the next call to Filter
		select {
		case <-b.done:
		}
		return
	}()

	return
}
