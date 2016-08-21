package thelm

import (
	"os"
	"runtime"
	"runtime/pprof"
)

type Profiler struct {
	cpuFile    string
	memoryFile string
}

func createProfileFile(outfile string) (fp *os.File, err error) {
	fp, err = os.Create(outfile)
	if err != nil {
		err = E.Annotate(err, "Could not create profile file", outfile)
	}
	return fp, err
}

func (p *Profiler) Setup(opts Options) (err error) {
	p.cpuFile = opts.Get("cpu-profile-file", "")
	p.memoryFile = opts.Get("memory-profile-file", "")

	if p.cpuFile != "" {
		var fp *os.File
		fp, err = createProfileFile(p.cpuFile)
		if err != nil {
			return
		}
		runtime.SetCPUProfileRate(1000)
		pprof.StartCPUProfile(fp)
	}

	return
}

func (p *Profiler) Close() (err error) {
	if p.cpuFile != "" {
		pprof.StopCPUProfile()
	}
	if p.memoryFile != "" {
		var fp *os.File
		fp, err = createProfileFile(p.memoryFile)
		if err != nil {
			return
		}
		defer fp.Close()
		pprof.WriteHeapProfile(fp)
	}
	return
}
