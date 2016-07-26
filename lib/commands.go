package thelm

import (
	"os/exec"
	"sync"
)

type Command struct {
	Out Buffer

	cmd   *exec.Cmd
	mutex sync.Mutex
}

func (c *Command) Finish() (err error) {
	c.mutex.Lock()
	if c.cmd != nil {
		c.cmd.Process.Kill()
	}
	c.mutex.Unlock()
	return
}

func (c *Command) Run(command string, args ...string) (err error) {
	c.Finish()
	c.Out.Reset()

	if command == "" {
		return
	}

	c.mutex.Lock()
	c.cmd = exec.Command(command, args...)
	c.cmd.Stdout = &c.Out
	c.cmd.Stderr = &c.Out
	err = c.cmd.Start()
	if err != nil {
		c.cmd = nil
	}
	c.mutex.Unlock()

	return
}
