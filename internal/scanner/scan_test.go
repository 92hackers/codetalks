/*

Test for the Scanner

*/

package scanner

import (
  "testing"
  "path/filepath"
  "fmt"

	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
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

func TestIsVCSDir(t *testing.T) {
  if !isVCSDir(".git") {
    t.Errorf(".git should not be a VCS directory")
  }

  if !isVCSDir(".svn") {
    t.Errorf(".svn should be a VCS directory")
  }

  if !isVCSDir(".hg") {
    t.Errorf(".hg should be a VCS directory")
  }

  if !isVCSDir(".bzr") {
    t.Errorf(".bzr should be a VCS directory")
  }

  if !isVCSDir(".cvs") {
    t.Errorf(".cvs should be a VCS directory")
  }
}

func TestScanEmptyCodebase(t *testing.T) {
  codeBase := filepath.Join("..", "..", "testdata/empty")
  rootDirs := []string{codeBase}
  Scan(rootDirs)
  fmt.Println(uniqueDirSet.Len())
  if uniqueDirSet.Len() != 0 {
    t.Errorf("Unique dirs set should be empty")
  }
  if len(language.AllLanguages) != 0 {
    t.Errorf("AllLanguages should be empty")
  }
  if len(file.AllCodeFiles) != 0 {
    t.Errorf("AllCodeFiles should be empty")
  }
}
