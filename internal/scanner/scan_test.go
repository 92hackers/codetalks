/*

Test for the Scanner

*/

package scanner

import (
  "testing"
  "path/filepath"

	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/utils"
	"github.com/92hackers/codetalks/internal/language"
)

func TestInit(t *testing.T) {
  utils.AssertNot(t, uniqueDirSet, nil)
  
  uniqueDirSet.Add("test")
  uniqueDirSet.Add("test2")

  utils.AssertEqual(t, uniqueDirSet.Len(), 2)
  utils.AssertEqual(t, uniqueDirSet.Contains("test"), true)

  uniqueDirSet.Remove("test2")

  utils.AssertEqual(t, uniqueDirSet.Len(), 1)
  utils.AssertEqual(t, uniqueDirSet.Contains("test2"), false)
  utils.AssertNot(t, uniqueDirSet.Contains("test333"), true)

  t.Cleanup(func() {
    uniqueDirSet = utils.NewSet() // Reset
  })
}

func TestConfig(t *testing.T) {
  t.Cleanup(func() {
    matchRegex = matchRegex[:0] // Reset
    ignoreRegex = ignoreRegex[:0] // Reset
  })

  utils.AssertEqual(t, len(matchRegex), 0)
  utils.AssertEqual(t, len(ignoreRegex), 0)

  Config("_test.go$", "")

  utils.AssertEqual(t, len(matchRegex), 1)
  utils.AssertEqual(t, len(ignoreRegex), 0)

  Config(" _test.go$", " vendor   ")

  utils.AssertEqual(t, len(matchRegex), 2)
  utils.AssertEqual(t, len(ignoreRegex), 1)

  utils.AssertEqual(t, ignoreRegex[0].MatchString("vendor"), true)
  utils.AssertEqual(t, ignoreRegex[0].MatchString("vendor/"), true)
  utils.AssertEqual(t, ignoreRegex[0].MatchString("vvvendor/"), true)

  utils.AssertEqual(t, matchRegex[1].MatchString("file_test.go"), true)
  utils.AssertEqual(t, matchRegex[1].MatchString("a/b/file_test.go"), true)
}

func TestIsVCSDir(t *testing.T) {
  vscDirs := []string{".git", ".svn", ".hg", ".bzr", ".cvs"}
  for _, dir := range vscDirs {
    utils.AssertEqual(t, isVCSDir(dir), true)
  }
  utils.AssertEqual(t, isVCSDir(".fakegit"), false)
  utils.AssertEqual(t, isVCSDir(".cvss"), false)
}

func TestScanEmptyCodebase(t *testing.T) {
  codeBase := filepath.Join("..", "..", "testdata/empty")
  rootDirs := []string{codeBase}
  Scan(rootDirs)
  utils.AssertEqual(t, uniqueDirSet.Len(), 0)
  utils.AssertEqual(t, len(language.AllLanguages), 0)
  utils.AssertEqual(t, len(file.AllCodeFiles), 0)
}

func TestScanSmallCodebase(t *testing.T) {
  codeBase := filepath.Join("..", "..", "testdata/small")
  rootDirs := []string{codeBase}
  Scan(rootDirs)
  utils.AssertEqual(t, uniqueDirSet.Len(), 2)
  utils.AssertEqual(t, len(language.AllLanguages), 2)
  utils.AssertEqual(t, len(language.AllLanguagesMap), 2)
  utils.AssertEqual(t, len(file.AllCodeFiles), 2)
}
