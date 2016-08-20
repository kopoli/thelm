package thelm

type ui struct {
	buffer []byte
}

// UiAbortedErr tells if user wanted to abort
var UiAbortedErr = E.New("User interface was aborted")

func Ui(opts Options, args []string) (ret string, err error) {

	return
}
