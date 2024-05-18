/**

time.go is a file that contains the time related functions and constants.

*/

package utils

import (
	"fmt"
	"time"
)

func TimeIt(fn func()) {
	start := time.Now()
	fn()
	fmt.Println("Time taken: ", time.Since(start))
}
