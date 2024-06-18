/**

Test utils

*/

package utils

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/fatih/color"
)

func AssertEqual(t *testing.T, actual any, expected any) {
	color.Set(color.FgRed)
	defer color.Unset()
	t.Helper()
	if expected != actual {
		Fail(t, expected, actual)
	}
}

func AssertNot(t *testing.T, actual any, expected any) {
	color.Set(color.FgRed)
	defer color.Unset()
	t.Helper()
	if expected == actual {
		Fail(t, expected, actual)
	}
}

func Fail(t *testing.T, actual any, expected any) {
	t.Helper()
	ErrorMsg("Error: Expected %v but got %v", expected, actual)
	t.FailNow()
}

func ErrorMsg(format string, a ...any) {
	color.NoColor = false
	color.Set(color.FgRed)
	defer color.Unset()
	color.Red(format, a...)
}

// CaptureStdout captures the output of a function that writes to stdout
// and returns the output content as a string
func CaptureStdout(fn func()) string {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
