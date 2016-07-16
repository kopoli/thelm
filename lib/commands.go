package thelm

import (
	"io"
	"os/exec"
	"sync"
)

type Command struct {
	Out Buffer

	running bool
	wg      sync.WaitGroup

	cmd *exec.Cmd
}

//TODO
// func CreateCommand(opt Options) (Command) {
// 	return Command{
// 		Out: make(TriggeringWriter),
// 	}
// }

func (c *Command) Setup(trigger func(), out io.Writer) {
	c.Out = Buffer{
		Trigger:     trigger,
		Passthrough: out,
	}
}

func (c *Command) Finish() (err error) {
	if c.running {
		err = c.cmd.Process.Kill()
		if err != nil {
			return
		}

		c.wg.Wait()
	}
	return
}

func (c *Command) Run(command string, args ...string) (err error) {
	c.Finish()
	c.Out.Reset()

	go func() {
		c.wg.Add(1)
		defer c.wg.Done()

		c.running = true

		c.cmd = exec.Command(command, args...)
		c.cmd.Stdout = &c.Out
		c.cmd.Stderr = &c.Out

		err = c.cmd.Run()

		c.running = false
	}()

	return
}
