/**

time.go is a file that contains the time related functions and constants.

*/

package utils

import (
	"fmt"
	"time"
)

func formatDuration(d time.Duration) string {
	scale := 100 * time.Second
	// look for the max scale that is smaller than d
	for scale > d {
		scale = scale / 10
	}
	return d.Round(scale / 100).String()
}

func TimeIt(fn func()) {
	start := time.Now()
	fn()
	fmt.Println("Time taken: ", formatDuration(time.Since(start)))
}

// AnalyzeTimeConsumed is a function that returns a function that can be used to analyze the time consumed.
// Usage: defer AnalyzeTimeConsumed()()
func AnalyzeTimeConsumed() func() {
	start := time.Now() // Closure
	return func() {
		fmt.Println("Analyze time consumed: ", formatDuration(time.Since(start)))
	}
}
