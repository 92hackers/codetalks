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
  if !s.Contains("apple") {
    t.Errorf("Set should contain apple")
  }

  // Remove an item from the set
  s.Remove("apple")

  // Check if an item is in the set
  if s.Contains("apple") {
    t.Errorf("Set should not contain apple")
  }

  // Check if an item is in the set
  if s.Contains("strawberry") {
    t.Errorf("Set should not contain strawberry")
  }
}
