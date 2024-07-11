/**

Testing for file package

*/

package file

import (
	"path/filepath"
	"testing"

	"github.com/92hackers/codetalks/internal/utils"
)

func clearState() {
	AllCodeFiles = AllCodeFiles[:0] // Reset
}

func TestNewCodeFile(t *testing.T) {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := NewCodeFile(f)

	utils.AssertEqual(t, len(AllCodeFiles), 1)
	utils.AssertEqual(t, file.Name, "hello.go")

	absPath, _ := filepath.Abs(f)
	utils.AssertEqual(t, file.Path, absPath)

	utils.AssertEqual(t, file.Language, "Go")
	utils.AssertEqual(t, file.FileExtension, ".go")
	utils.AssertEqual(t, file.FileType, CODE_FILE)
	utils.AssertNot(t, file.LastModifiedAt, uint64(0))
	utils.AssertNot(t, file.Directory, "")

	utils.AssertEqual(t, file.Size, uint64(0))
	utils.AssertEqual(t, file.Content, "")
	utils.AssertEqual(t, file.CodeCount, uint32(0))
	utils.AssertEqual(t, file.CommentCount, uint32(0))
	utils.AssertEqual(t, file.BlankCount, uint32(0))
	utils.AssertEqual(t, file.TotalLines, uint32(0))

	t.Cleanup(clearState)
}

func TestNewCodeFileWithInvalidPath(t *testing.T) {
	file, err := NewCodeFile("invalid_path")
	utils.AssertNot(t, err, nil)
	utils.AssertEqual(t, file, (*CodeFile)(nil))

	t.Cleanup(clearState)
}

func TestAnalyze(t *testing.T) {
	f := filepath.Join("..", "..", "testdata/small/hello.go")
	file, _ := NewCodeFile(f)

	ctx, cancel := utils.WithTimeoutCtxSeconds(1)
	defer cancel()
	file.Analyze(ctx)

	utils.AssertEqual(t, file.CodeCount, uint32(7))
	utils.AssertEqual(t, file.CommentCount, uint32(4))
	utils.AssertEqual(t, file.BlankCount, uint32(3))
	utils.AssertEqual(t, file.TotalLines, uint32(14))
	utils.AssertNot(t, file.Size, uint64(0))

	t.Cleanup(clearState)
}
