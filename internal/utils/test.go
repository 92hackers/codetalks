/**

Test utils

*/

package utils

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func AssertEqual(t *testing.T, actual any, expected any) {
	t.Helper()
	if expected != actual {
		t.Errorf("❌ Expected %v but got %v", expected, actual)
	}
}

func AssertNot(t *testing.T, actual any, expected any) {
	t.Helper()
	if expected == actual {
		t.Errorf("❌ Expected %v but got %v", expected, actual)
	}
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
