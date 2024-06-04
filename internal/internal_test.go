/**

Test for internal common code.

*/

package internal

import (
	"github.com/92hackers/codetalks/internal/utils"
	"testing"
)

func TestNewLanguageDef(t *testing.T) {
	golang := newLangDef("golang", []string{"//"}, [][]string{{"/*", "*/"}})
	utils.AssertEqual(t, golang.Name, "golang")
	utils.AssertEqual(t, golang.LineComments[0], "//")
	utils.AssertEqual(t, len(golang.LineComments), 1)
	utils.AssertEqual(t, len(golang.BlockComments), 1)
	utils.AssertEqual(t, golang.BlockComments[0][0], "/*")
	utils.AssertEqual(t, golang.BlockComments[0][1], "*/")
}

func TestNewLanguageDefWithEmptyBlockComments(t *testing.T) {
	golang := newLangDef("golang", []string{"#"}, [][]string{})
	utils.AssertEqual(t, golang.Name, "golang")
	utils.AssertEqual(t, golang.LineComments[0], "#")
	utils.AssertEqual(t, len(golang.LineComments), 1)
	utils.AssertEqual(t, len(golang.BlockComments), 0)
}

func TestNewLanguageDefWithEmptyLineComments(t *testing.T) {
	golang := newLangDef("golang", []string{}, [][]string{{"/*", "*/"}})
	utils.AssertEqual(t, golang.Name, "golang")
	utils.AssertEqual(t, len(golang.LineComments), 0)
	utils.AssertEqual(t, len(golang.BlockComments), 1)
	utils.AssertEqual(t, golang.BlockComments[0][0], "/*")
	utils.AssertEqual(t, golang.BlockComments[0][1], "*/")
}

func TestDebugMode(t *testing.T) {
	utils.AssertEqual(t, IsDebugEnabled, false)
	IsDebugEnabled = true
	utils.AssertEqual(t, IsDebugEnabled, true)
}

func TestSupportedLanguages(t *testing.T) {
	utils.AssertNot(t, len(SupportedLanguages), 0)
	exts := []string{
		".s", ".asm", ".S", ".c", ".h", ".cpp", ".hpp", ".hh",
		".go", ".java", ".js", ".ts", ".kt", ".rs", ".py", ".rb", ".php",
		".sh", ".pl", ".lua", ".swift", ".cs", ".r", ".scala", ".md", ".txt",
		".zsh", ".bash", ".cjs", ".cts", ".jsx", ".mts", ".mjs", ".tsx",
	}
	for _, ext := range exts {
		if SupportedLanguages[ext] == nil {
			t.Errorf("❌ Language definition for %s is nil", ext)
		}
	}
	// Test for a language that is not supported
	if SupportedLanguages[".xx"] != nil {
		t.Errorf("❌ Language definition for .xx should be nil")
	}
}
