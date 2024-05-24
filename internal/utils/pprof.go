/**
pprof util
*/

package utils

import (
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

// You can view the cpu profile file with `go tool pprof codetalks-cpu.prof`
// or visualize it on website like https://www.speedscope.app/
func StartCPUProfile() func() {
	f, err := os.Create("codetalks-cpu.prof")
	if err != nil {
		fmt.Println("Create cpu profile file failed with error: ", err)
		return nil
	}
	pprof.StartCPUProfile(f)
	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}

// You can view the trace file with `go tool trace codetalks-trace.prof`
func StartTrace() func() {
	f, err := os.Create("codetalks-trace.prof")
	if err != nil {
		fmt.Println("Create trace file failed with error: ", err)
		return nil
	}
	trace.Start(f)
	return func() {
		trace.Stop()
		f.Close()
	}
}

// You can view the memory profile file with `go tool pprof codetalks-mem.prof`
// or visualize it on website like https://www.speedscope.app/
func StartMemoryProfile() func() {
	f, err := os.Create("codetalks-mem.prof")
	if err != nil {
		fmt.Println("Create memory profile file failed with error: ", err)
		return nil
	}
	return func() {
		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Println("Write memory profile failed with error: ", err)
		}
		f.Close()
	}
}
