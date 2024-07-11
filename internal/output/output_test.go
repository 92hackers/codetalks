/**

Test for the output package

*/

package output

import (
	"path/filepath"
	"testing"

	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
	"github.com/92hackers/codetalks/internal/utils"
)

func generateData() {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := file.NewCodeFile(f)
	ctx, cancel := utils.WithTimeoutCtxSeconds(1)
	defer cancel()
	file.Analyze(ctx)
	lang := language.AddLanguage(".go", file)
	lang.CountCodeFileStats(file)
	language.AggreateStats()
}

func TestOutputCliTableHeader(t *testing.T) {
	output := utils.CaptureStdout(renderHeader)
	expected := `===============================================================================
| Language     | Files     | Total      | Comments   | Blanks    | Code       |
===============================================================================
`
	utils.AssertEqual(t, len(output), len(expected))
	utils.AssertEqual(t, output, expected)
}

func clearState() {
	language.AllLanguages = language.AllLanguages[:0]                      // Reset
	language.AllLanguagesMap = make(map[string]uint)                       // Reset
	language.AllLanguageAggregateStats = language.LanguageAggregateStats{} // Reset
	file.AllCodeFiles = file.AllCodeFiles[:0]                              // Reset
}

func TestOutputCliTableData(t *testing.T) {
	generateData()
	output := utils.CaptureStdout(renderData)
	expected := `| Go           | 1         | 14         | 4          | 3         | 7          |
`
	utils.AssertEqual(t, len(output), len(expected))
	utils.AssertEqual(t, output, expected)

	t.Cleanup(clearState)
}

func TestOutputCliTableFooterEmptyData(t *testing.T) {
	output := utils.CaptureStdout(renderFooter)
	expected := `| Total        | 0         | 0          | 0          | 0         | 0          |
===============================================================================
`
	utils.AssertEqual(t, len(output), len(expected))
	utils.AssertEqual(t, output, expected)
}

func TestOutputCliTableWithoutData(t *testing.T) {
	output := utils.CaptureStdout(OutputCliTable)
	expected := `===============================================================================
| Language     | Files     | Total      | Comments   | Blanks    | Code       |
===============================================================================
| Total        | 0         | 0          | 0          | 0         | 0          |
===============================================================================
`
	utils.AssertEqual(t, len(output), len(expected))
	utils.AssertEqual(t, output, expected)
}

func TestOutputCliTable(t *testing.T) {
	generateData()
	output := utils.CaptureStdout(OutputCliTable)
	expected := `===============================================================================
| Language     | Files     | Total      | Comments   | Blanks    | Code       |
===============================================================================
| Go           | 1         | 14         | 4          | 3         | 7          |
===============================================================================
| Total        | 1         | 14         | 4          | 3         | 7          |
===============================================================================
`
	utils.AssertEqual(t, len(output), len(expected))
	utils.AssertEqual(t, output, expected)

	t.Cleanup(clearState)
}

func TestOutputJSONEmptyData(t *testing.T) {
	t.Skip()
	output := utils.CaptureStdout(OutputJSON)
	expected := `{
  "languages": [],
  "total": {
    "files": 0,
    "total": 0,
    "comments": 0,
    "blanks": 0,
    "code": 0
  }
}
`
	utils.AssertEqual(t, len(output), len(expected))
	utils.AssertEqual(t, output, expected)
}

func TestOutputJSON(t *testing.T) {
	t.Skip()
	generateData()
	output := utils.CaptureStdout(OutputJSON)
	expected := `{
  "languages": [
    {
      "name": "Go",
      "files": 1,
      "total": 14,
      "comments": 4,
      "blanks": 3,
      "code": 7
    }
  ],
  "total": {
    "files": 1,
    "total": 14,
    "comments": 4,
    "blanks": 3,
    "code": 7
  }
}
`
	utils.AssertEqual(t, len(output), len(expected))
	utils.AssertEqual(t, output, expected)

	t.Cleanup(clearState)
}
