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
