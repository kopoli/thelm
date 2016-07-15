package thelm

import (
	"bytes"
	"io"
	"os/exec"
	"sync"
)

type TriggeringWriter struct {
	Trigger func()
	Writer  io.Writer
	Count   int
}

func (tw *TriggeringWriter) Write(p []byte) (n int, err error) {
	tw.Count = tw.Count + bytes.Count(p, []byte("\n"))
	tw.Trigger()
	return tw.Writer.Write(p)
}

type Command struct {
	Out TriggeringWriter

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

	c.Out.Count = 0

	go func() {
		c.wg.Add(1)
		defer c.wg.Done()

		c.running = true

		c.cmd = exec.Command(command, args...)
		c.cmd.Stdout = &c.Out
		c.cmd.Stderr = &c.Out

		err = c.cmd.Run()

		// if err != nil {
		// 	// fmt.Println("running command failed:", err)
		// }
		c.running = false
	}()

	return
}
