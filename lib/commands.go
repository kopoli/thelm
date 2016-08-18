package thelm

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
)

type Command struct {
	cmd   *exec.Cmd
	mutex sync.Mutex

	// Callback that provides data out
	Sync func([]byte) error

	lines    int
	MaxLines int
}

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
	n = len(p)
	c.Sync(p)
	return
}

// Finish makes sure the process has stopped
func (c *Command) Finish() (err error) {
	c.mutex.Lock()
	if c.cmd != nil {
		c.cmd.Process.Kill()
	}
	c.mutex.Unlock()
	return
}

// Run given cocmmand with args
func (c *Command) Run(command string, args ...string) (err error) {
	c.Finish()

	if command == "" {
		return
	}

	c.mutex.Lock()
	c.cmd = exec.Command(command, args...)
	c.cmd.Stdout = c
	c.cmd.Stderr = c
	c.lines = 0
	err = c.cmd.Start()
	if err != nil {
		c.cmd = nil
		c.Sync([]byte{})
	}
	c.mutex.Unlock()

	return
}
