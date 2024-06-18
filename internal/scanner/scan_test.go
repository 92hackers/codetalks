/*

Test for the Scanner

*/

package scanner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
	"github.com/92hackers/codetalks/internal/utils"
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
		matchRegex = matchRegex[:0]   // Reset
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

func clearState() {
	uniqueDirSet = utils.NewSet()                     // Reset
	language.AllLanguages = language.AllLanguages[:0] // Reset
	language.AllLanguagesMap = make(map[string]uint)  // Reset
	file.AllCodeFiles = file.AllCodeFiles[:0]         // Reset
	matchRegex = matchRegex[:0]                       // Reset
	ignoreRegex = ignoreRegex[:0]                     // Reset
}

func TestScanSmallCodebase(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase}
	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 2)
	utils.AssertEqual(t, len(language.AllLanguages), 2)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 2)
	utils.AssertEqual(t, len(file.AllCodeFiles), 2)

	t.Cleanup(clearState)
}

func TestScanSmallCodebaseMatchOption(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase}

	Config(" go$ ", "")
	utils.AssertEqual(t, len(matchRegex), 1)
	utils.AssertEqual(t, len(ignoreRegex), 0)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 1)
	utils.AssertEqual(t, len(language.AllLanguages), 1)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 1)
	utils.AssertEqual(t, len(file.AllCodeFiles), 1)

	t.Cleanup(clearState)
}

func TestScanSmallCodebaseIgnoreOptionNoWork(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase}

	Config("", "vendor")
	utils.AssertEqual(t, len(matchRegex), 0)
	utils.AssertEqual(t, len(ignoreRegex), 1)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 2)
	utils.AssertEqual(t, len(language.AllLanguages), 2)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 2)
	utils.AssertEqual(t, len(file.AllCodeFiles), 2)

	t.Cleanup(clearState)
}

func TestScanSmallCodebaseIgnoreOption(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase}

	Config("", ".+.py$")
	utils.AssertEqual(t, len(matchRegex), 0)
	utils.AssertEqual(t, len(ignoreRegex), 1)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 1)
	utils.AssertEqual(t, len(language.AllLanguages), 1)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 1)
	utils.AssertEqual(t, len(file.AllCodeFiles), 1)

	t.Cleanup(clearState)
}

func TestScanSmallCodebaseSameMatchAndIgnoreOptions(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase}

	Config(".+.py$", ".+.py$")
	utils.AssertEqual(t, len(matchRegex), 1)
	utils.AssertEqual(t, len(ignoreRegex), 1)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 0)
	utils.AssertEqual(t, len(language.AllLanguages), 0)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 0)
	utils.AssertEqual(t, len(file.AllCodeFiles), 0)

	t.Cleanup(clearState)
}

func TestScanSmallCodebaseIgnoreAll(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase}

	Config("", ".+")
	utils.AssertEqual(t, len(matchRegex), 0)
	utils.AssertEqual(t, len(ignoreRegex), 1)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 0)
	utils.AssertEqual(t, len(language.AllLanguages), 0)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 0)
	utils.AssertEqual(t, len(file.AllCodeFiles), 0)

	t.Cleanup(clearState)
}

func TestScanSmallCodebaseMatchAndIgnoreOption(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase}

	Config(" go$ ", ".+.py$  ")
	utils.AssertEqual(t, len(matchRegex), 1)
	utils.AssertEqual(t, len(ignoreRegex), 1)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 1)
	utils.AssertEqual(t, len(language.AllLanguages), 1)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 1)
	utils.AssertEqual(t, len(file.AllCodeFiles), 1)

	t.Cleanup(clearState)
}

func TestScanSmallCodebaseDuplicateRootDirs(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/small")
	rootDirs := []string{codeBase, codeBase, codeBase}

	Config(" go$ ", ".+.py$  ")
	utils.AssertEqual(t, len(matchRegex), 1)
	utils.AssertEqual(t, len(ignoreRegex), 1)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 1)
	utils.AssertEqual(t, len(language.AllLanguages), 1)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 1)
	utils.AssertEqual(t, len(file.AllCodeFiles), 1)

	t.Cleanup(clearState)
}

func TestScanUnsupportedFiles(t *testing.T) {
	codeBase := filepath.Join("..", "..", "testdata/unsupported")
	rootDirs := []string{codeBase}

	Config(" go$ ", ".+.py$  ")
	utils.AssertEqual(t, len(matchRegex), 1)
	utils.AssertEqual(t, len(ignoreRegex), 1)

	Scan(rootDirs)
	utils.AssertEqual(t, uniqueDirSet.Len(), 0)
	utils.AssertEqual(t, len(language.AllLanguages), 0)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 0)
	utils.AssertEqual(t, len(file.AllCodeFiles), 0)

	t.Cleanup(clearState)
}

func mockCodebase(t *testing.T) (rootDirs []string) {
	t.Helper()
	codeBase := t.TempDir()
	rootDirs = append(rootDirs, codeBase)

	createFile := func(name, content string) {
		fd, _ := os.Create(filepath.Join(codeBase, name))
		defer fd.Close()
		fd.WriteString(content)
	}

	// Create a few files and dirs
	createFile("main.go", "package main\n\nfunc main() {\n\tprintln(\"Hello, World!\")\n}\n")
	createFile("main.py", "print('Hello, World!')\n")
	createFile("main.js", "console.log('Hello, World!')\n")
	createFile(".gitignore", "*.py\n*.js\nvendor\n")
	os.Mkdir(filepath.Join(codeBase, "vendor"), 0755)
	createFile("vendor/main.go", "package main\n\nfunc main() {\n\tprintln(\"Hello, World!\")\n}\n")
	createFile("vendor/main.sh", "echo 'Hello, World!'\n")

	return
}

func TestScannerRespectGitIgnore(t *testing.T) {
	rootDirs := mockCodebase(t)
	Scan(rootDirs)

	utils.AssertEqual(t, uniqueDirSet.Len(), 1)
	utils.AssertEqual(t, len(language.AllLanguages), 1)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 1)
	utils.AssertEqual(t, len(file.AllCodeFiles), 1)

	t.Cleanup(clearState)
}

func TestScannerRespectGitIgnoreWithMatch(t *testing.T) {
	rootDirs := mockCodebase(t)
	Config(" .go$", "")
	Scan(rootDirs)

	utils.AssertEqual(t, uniqueDirSet.Len(), 1)
	utils.AssertEqual(t, len(language.AllLanguages), 1)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 1)
	utils.AssertEqual(t, len(file.AllCodeFiles), 1)

	t.Cleanup(clearState)
}

func TestScannerRespectGitIgnoreWithIgnore(t *testing.T) {
	rootDirs := mockCodebase(t)
	Config("", "vendor .js$")
	Scan(rootDirs)

	utils.AssertEqual(t, uniqueDirSet.Len(), 1)
	utils.AssertEqual(t, len(language.AllLanguages), 1)
	utils.AssertEqual(t, len(language.AllLanguagesMap), 1)
	utils.AssertEqual(t, len(file.AllCodeFiles), 1)

	t.Cleanup(clearState)
}

func TestIsSpecifiedDepthDirs(t *testing.T) {
	type PathData struct {
		path  string // The path will be relative to a root directory in source code.
		depth int
	}
	truePaths := []PathData{
		{"test", 1},
		{"/test", 1},
		{"/test/", 1},
		{"/test//", 1},
		{"/test/a", 2},
		{"/test/a/b", 3},
		{"/test/a/b/", 3},
	}
	falsePaths := []PathData{
		{"test", 0},
		{"/test", 0},
		{"/test/", 2},
		{"/test//", 2},
		{"/test/a", 1},
		{"/test/a/b", 2},
	}

	for _, pathData := range truePaths {
		t.Run("truepath---"+pathData.path, func(t *testing.T) {
			utils.AssertEqual(t, isSpecifiedDepthDirs(pathData.path, pathData.depth), true)
		})
	}
	for _, pathData := range falsePaths {
		t.Run("falsepath---"+pathData.path, func(t *testing.T) {
			utils.AssertEqual(t, isSpecifiedDepthDirs(pathData.path, pathData.depth), false)
		})
	}
}
