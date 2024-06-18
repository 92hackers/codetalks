/**

Testing for language package

*/

package language

import (
	"go.uber.org/goleak"
	"path/filepath"
	"testing"

	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/utils"
)

// TestMain is the entry point for the test
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m) // Check for goroutine leaks
}

func clearState() {
	AllLanguages = AllLanguages[:0]                      // Reset
	AllLanguagesMap = make(map[string]uint)              // Reset
	AllLanguageAggregateStats = LanguageAggregateStats{} // Reset
	file.AllCodeFiles = file.AllCodeFiles[:0]            // Reset
}

func TestAddLanguage(t *testing.T) {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := file.NewCodeFile(f)
	file.Analyze()
	lang := AddLanguage(".go", file)

	utils.AssertEqual(t, len(AllLanguages), 1)
	utils.AssertEqual(t, len(AllLanguagesMap), 1)
	utils.AssertEqual(t, len(lang.CodeFiles), 1)
	utils.AssertEqual(t, lang.FileCount, uint32(1))
	utils.AssertEqual(t, lang.Name, "Go")
	utils.AssertEqual(t, lang.FileExtension, ".go")

	lang.CountCodeFileStats(file)

	utils.AssertEqual(t, lang.CodeCount, uint32(7))
	utils.AssertEqual(t, lang.CommentCount, uint32(4))
	utils.AssertEqual(t, lang.BlankCount, uint32(3))
	utils.AssertEqual(t, lang.TotalLines, uint32(14))

	t.Cleanup(clearState)
}

func TestGetLanguage(t *testing.T) {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := file.NewCodeFile(f)
	file.Analyze()
	AddLanguage(".go", file)

	lang := GetLanguage(".go")

	utils.AssertEqual(t, len(AllLanguages), 1)
	utils.AssertEqual(t, lang.Name, "Go")
	utils.AssertEqual(t, lang.FileExtension, ".go")
	utils.AssertEqual(t, len(lang.CodeFiles), 1)

	t.Cleanup(clearState)
}

func TestGetLanguageNonExisted(t *testing.T) {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := file.NewCodeFile(f)
	file.Analyze()
	AddLanguage(".go", file)

	lang := GetLanguage(".pyy")

	if lang != nil {
		t.Errorf("Expected nil, got %v", lang)
	}

	t.Cleanup(clearState)
}

func TestAggreateStats(t *testing.T) {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := file.NewCodeFile(f)
	file.Analyze()
	lang := AddLanguage(".go", file)

	lang.CountCodeFileStats(file)

	AggreateStats()

	utils.AssertEqual(t, AllLanguageAggregateStats.TotalCode, uint32(7))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalComment, uint32(4))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalBlank, uint32(3))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalLines, uint32(14))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalFiles, uint32(1))

	t.Cleanup(clearState)
}

func TestAnalyzeAllLanguages(t *testing.T) {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := file.NewCodeFile(f)
	lang := AddLanguage(".go", file)

	utils.AssertEqual(t, lang.CodeCount, uint32(0))
	utils.AssertEqual(t, lang.CommentCount, uint32(0))
	utils.AssertEqual(t, lang.BlankCount, uint32(0))
	utils.AssertEqual(t, lang.TotalLines, uint32(0))

	AnalyzeAllLanguages()

	utils.AssertEqual(t, lang.CodeCount, uint32(7))
	utils.AssertEqual(t, lang.CommentCount, uint32(4))
	utils.AssertEqual(t, lang.BlankCount, uint32(3))
	utils.AssertEqual(t, lang.TotalLines, uint32(14))

	utils.AssertEqual(t, len(AllLanguages), 1)
	utils.AssertEqual(t, len(AllLanguagesMap), 1)
	utils.AssertEqual(t, len(AllLanguages[0].CodeFiles), 1)

	utils.AssertEqual(t, AllLanguageAggregateStats.TotalCode, uint32(7))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalComment, uint32(4))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalBlank, uint32(3))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalLines, uint32(14))
	utils.AssertEqual(t, AllLanguageAggregateStats.TotalFiles, uint32(1))

	t.Cleanup(clearState)
}
