package thelm

import (
	"fmt"
	"os"
)

const thelmEnv = "_THELM_RUNNING_"

func CheckSelfRunning(opts Options) (err error) {
	running := os.Getenv(thelmEnv)
	if running != "" {
		progName := opts.Get("program-name", "thelm")
		return E.New(fmt.Sprintf("%s detected running inside %s", progName, progName))
	}

	return os.Setenv(thelmEnv, "t")
}
