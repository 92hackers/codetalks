/**

Test utils

*/

package utils

import (
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
