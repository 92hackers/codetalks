/*

Test for the Scanner

*/

package scanner

import (
  "testing"
  "fmt"
)

func TestInit(t *testing.T) {
  if uniqueDirSet == nil {
    t.Errorf("uniqueDirSet should not be nil")
  }
  
  uniqueDirSet.Add("test")
  uniqueDirSet.Add("test2")

  if !uniqueDirSet.Contains("test") {
    t.Errorf("uniqueDirSet should contain test")
  }

  uniqueDirSet.Remove("test2")

  if uniqueDirSet.Contains("test2") {
    t.Errorf("uniqueDirSet should not contain test2")
  }

  if uniqueDirSet.Contains("test3") {
    t.Errorf("uniqueDirSet should not contain test3")
  }
}

func TestConfig(t *testing.T) {

  t.Cleanup(func() {
    fmt.Println("Cleanup...")
    matchRegex = matchRegex[:0] // Reset
    ignoreRegex = ignoreRegex[:0] // Reset
  })

  if (len(matchRegex) != 0) {
    t.Errorf("matchRegex should be empty")
  }
  if (len(ignoreRegex) != 0) {
    t.Errorf("ignoreRegex should be empty")
  }

  Config("_test.go$", "")

  if len(matchRegex) != 1 {
    t.Errorf("matchRegex should have 1 element")
  }

  if len(ignoreRegex) != 0 {
    t.Errorf("ignoreRegex should have be empty")
  }

  Config(" _test.go$", " vendor   ")

  if len(matchRegex) != 2 {
    t.Errorf("matchRegex should have 2 element")
  }

  if len(ignoreRegex) != 1 {
    t.Errorf("ignoreRegex should have 1 element")
  }

  if !ignoreRegex[0].MatchString("vendor") {
    t.Errorf("ignoreRegex should match vendor")
  }

  if !matchRegex[1].MatchString("file_test.go") {
    t.Errorf("matchRegex should match test.go")
  }

  if !matchRegex[1].MatchString("a/b/file_test.go") {
    t.Errorf("matchRegex should match test.go")
  }
}

