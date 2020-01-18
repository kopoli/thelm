package thelm

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

type AsyncSource interface {
	Run(...string) error
	Wait()
	Finish() error

	IsOneShot() bool

	// Set callback that provides data out
	SetOutput(io.Writer)
}

type Source struct {
	mutex  sync.Mutex
	output io.Writer
}

func (s *Source) SetOutput(p io.Writer) {
	s.output = p
}

type SourceReader struct {
	Input     io.Reader
	wg        sync.WaitGroup
	isRead    bool
	sourceErr error

	Source
}

var _ AsyncSource = (*SourceReader)(nil)

func (r *SourceReader) IsOneShot() bool {
	return true
}

func (r *SourceReader) Run(...string) error {
	r.Wait()

	r.mutex.Lock()
	if r.isRead {
		r.mutex.Unlock()
		return nil
	}
	r.mutex.Unlock()
	r.wg.Add(1)

	go func() {
		r.mutex.Lock()
		_, r.sourceErr = io.Copy(r.output, r.Input)
		r.isRead = true
		r.wg.Done()
		r.mutex.Unlock()
	}()
	return nil
}

func (r *SourceReader) Finish() (err error) {
	r.Wait()
	r.mutex.Lock()
	err = r.sourceErr
	r.mutex.Unlock()
	return
}

func (r *SourceReader) Wait() {
	r.wg.Wait()
}

type SourceFile struct {
	FileName string

	SourceReader
}

func (f *SourceFile) Run(args ...string) (err error) {
	var file *os.File
	file, err = os.Open(f.FileName)
	if err != nil {
		err = E.Annotate(err, "Could not open file to read")
		return
	}
	f.Input = file
	err = f.SourceReader.Run(args...)
	return
}

func (f *SourceFile) Wait() {
	f.SourceReader.Wait()
	if f.Input != nil {
		f.Input.(*os.File).Close()
	}
}

type Command struct {
	cmd      *exec.Cmd
	lines    int
	MaxLines int

	Source
}

var _ io.Writer = (*Command)(nil)
var _ AsyncSource = (*Command)(nil)

// Data will be written to the internal buffer from another process
func (c *Command) Write(p []byte) (n int, err error) {
	lines := bytes.Count(p, []byte("\n"))
	if c.MaxLines != 0 && c.lines+lines > c.MaxLines {
		linedata := bytes.SplitN(p, []byte("\n"), -1)
		p = bytes.Join(linedata[0:c.MaxLines-c.lines], []byte("\n"))
		p = append(p, []byte(fmt.Sprintf("\n-- command output cut off at %d lines --\n", c.MaxLines))...)
		lines = c.MaxLines - c.lines
		c.Finish()
	}

	c.lines += lines
	return c.output.Write(p)
}

// Finish makes sure the process has stopped
func (c *Command) Finish() (err error) {
	c.mutex.Lock()
	if c.cmd != nil {
		_ = c.cmd.Process.Kill()
		_ = c.cmd.Wait()
	}
	c.mutex.Unlock()
	return
}

// Wait waits for the command to complete
func (c *Command) Wait() {
	c.mutex.Lock()
	if c.cmd != nil {
		c.cmd.Wait()
	}
	c.mutex.Unlock()
}

// Run given command with args
func (c *Command) Run(args ...string) (err error) {
	c.Finish()

	if len(args) == 0 || args[0] == "" {
		return
	}

	c.mutex.Lock()
	c.cmd = exec.Command(args[0], args[1:]...)
	c.cmd.Stdout = c
	c.cmd.Stderr = c
	c.lines = 0
	err = c.cmd.Start()
	if err != nil {
		c.cmd = nil
	}
	c.mutex.Unlock()

	return
}

func (c *Command) IsOneShot() bool {
	return false
}
