package bin

import (
	"runtime"
)

const MIN_PROCS = 5

func InitProcs() {
	if runtime.GOMAXPROCS(-1) < MIN_PROCS {
		runtime.GOMAXPROCS(MIN_PROCS)
	}
}

func NumProcs() int {
	return runtime.GOMAXPROCS(-1)
}
