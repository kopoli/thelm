package thelm

import (
	"bytes"
	"io"
	"regexp"
	"sync"
)

// Buffer stores data written to it through its io.Writer. It also writes the
// data to Passthrough. Each time Write is called it counts the lines to Count
// and runs the Trigger function
type Buffer struct {
	Passthrough io.Writer
	Count       int
	AfterWrite  func()
	OnStart     func() error
	data        []byte
	pos         int
	disabled    bool
	mutex       sync.Mutex
}

func (b *Buffer) DisableWriting() {
	b.disabled = true
}

func (b *Buffer) Reset() {
	b.OnStart()

	b.Count = 0
	b.data = []byte{}
	b.pos = 0
	b.disabled = false
}

func (b *Buffer) Sync() (err error) {
	b.mutex.Lock()
	_, err = b.Passthrough.Write(b.data[b.pos:len(b.data)])
	b.pos = len(b.data)
	b.mutex.Unlock()
	return
}

// Write data into the buffer and through to Passthrough.
func (b *Buffer) Write(p []byte) (n int, err error) {
	if b.disabled {
		return
	}

	b.Count += bytes.Count(p, []byte("\n"))
	b.mutex.Lock()
	b.data = append(b.data, p...)
	n = len(p)
	b.mutex.Unlock()
	b.AfterWrite()
	return
}

// Filter the current Buffer with the regexp and write output to Passthrough.
func (b *Buffer) Filter(regex string) (err error) {
	if b.disabled {
		return
	}

	re, err := regexp.Compile("(?i)" + AsRelaxedRegexp(regex))
	if err != nil {
		return
	}

	b.OnStart()
	b.Count = 0

	for _, line := range bytes.Split(b.data, []byte("\n")) {
		if re.Match(line) {
			b.Count += 1
			_, err = b.Passthrough.Write(append(line, '\n'))
			if err != nil {
				return
			}
			b.AfterWrite()
		}
	}

	return
}

func (b *Buffer) RestoreFiltering() {
	b.Count = len(b.data)
}
