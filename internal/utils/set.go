/**

A set data structure.

*/

package utils

import (
	"sync"
)

// Define an empty struct to use as the dummy value
type void struct{}

// Define a set as a map with keys of type string and values of type void
type Set struct {
	m     map[string]void
	mutex sync.RWMutex
}

// Helper function to create a new set
func NewSet() *Set {
	return &Set{m: make(map[string]void)}
}

// Helper function to add an item to the set
func (s *Set) Add(item string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[item] = void{}
}

// Helper function to remove an item from the set
func (s *Set) Remove(item string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.m, item)
}

// Helper function to check if an item is in the set
func (s *Set) Contains(item string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	_, ok := s.m[item]
	return ok
}

// Helper function to get the number of items in the set
func (s *Set) Len() int {
  s.mutex.RLock()
  defer s.mutex.RUnlock()
  return len(s.m)
}

// Usage:
//
// func main() {
//   // Create a new set
//   s := utils.NewSet()
//
//   // Add some items to the set
//   s.add("apple")
//   s.add("banana")
//   s.add("cherry")
//
//   // Check if an item is in the set
//   fmt.Println("Contains apple:", s.contains("apple"))
//   fmt.Println("Contains grape:", s.contains("grape"))
//
//   // Remove an item from the set
//   s.remove("banana")
//
//   // Check if the item was removed
//   fmt.Println("Contains banana:", s.contains("banana"))
// }
