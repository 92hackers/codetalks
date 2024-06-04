/**

Test for the Set data structure.

*/

package utils

import (
	"testing"
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
