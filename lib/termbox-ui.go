package thelm

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
)

// UiAbortedErr tells if user wanted to abort
var UiAbortedErr = E.New("User interface was aborted")

func Ui(opts Options, args []string) (ret string, err error) {

	var view UiView

	// Termbox setup
	err = termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	// Testing
	view.SetStatusLine(" Thelm testi ")
	fmt.Fprintln(&view, "viewtesti")
	fmt.Fprintln(&view, "Tahan toiselle riville")
	fmt.Fprintln(&view, "Kolmas rivi")
	view.Flush()

	// Main loop
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			}
		case termbox.EventError:
			err = ev.Err
			return
		}
	}

	return
}
