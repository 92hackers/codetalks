/**

Test for the Set data structure.

*/

package utils

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	// Create a new set
	s := NewSet()

	// Add some items to the set
	s.Add("apple")
	s.Add("banana")
	s.Add("cherry")

	// Check if an item is in the set
	AssertEqual(t, s.Len(), 3)

	AssertEqual(t, s.Contains("apple"), true)
	AssertEqual(t, s.Contains("banana"), true)
	AssertEqual(t, s.Contains("cherry"), true)

	// Remove an item from the set
	s.Remove("apple")

	// Check if an item is in the set
	AssertEqual(t, s.Len(), 2)
	AssertNot(t, s.Contains("apple"), true)

	// Check if an item is in the set
	AssertNot(t, s.Contains("strawberry"), true)
}

func TestTimeIt(t *testing.T) {
	// TimeIt
	timeTakenStr := CaptureStdout(func() {
		TimeIt(func() {
			time.Sleep(100 * time.Millisecond)
		})
	})
	AssertEqual(t, timeTakenStr, "Time taken: 100ms\n")
}

func TestFormatDuration(t *testing.T) {
	// FormatDuration
	AssertEqual(t, formatDuration(1), "1ns")
	AssertEqual(t, formatDuration(1000), "1µs")
	AssertEqual(t, formatDuration(1000000), "1ms")
	AssertEqual(t, formatDuration(1000000000), "1s")
	AssertEqual(t, formatDuration(3578970000), "3.58s")
	AssertEqual(t, formatDuration(1000000000000), "16m40s")
}

func TestAnalyzeTimeConsumed(t *testing.T) {
	// AnalyzeTimeConsumed
	fn := AnalyzeTimeConsumed()
	time.Sleep(125 * time.Millisecond)
	analyzeTimeConsumedStr := CaptureStdout(fn)
	AssertEqual(t, analyzeTimeConsumedStr, "Analyze time consumed: 125ms\n")
}

func TestWithTimeoutCtxSeconds(t *testing.T) {
	// WithTimeoutCtxSeconds
	ctx, cancel := WithTimeoutCtxSeconds(1)
	defer cancel()

	select {
	case <-ctx.Done():
		AssertEqual(t, ctx.Err().Error(), "context deadline exceeded")
	case <-time.After(2 * time.Second):
		t.Error("timeout not working")
	}
}

func TestWithTimeoutCtxMilliSeconds(t *testing.T) {
	// WithTimeoutCtxMilliSeconds
	ctx, cancel := WithTimeoutCtxMilliSeconds(1000)
	defer cancel()

	select {
	case <-ctx.Done():
		AssertEqual(t, ctx.Err().Error(), "context deadline exceeded")
	case <-time.After(2 * time.Second):
		t.Error("timeout not working")
	}
}
