/**
Code file
*/

package file

import (
	"log"
	"os"
	"path/filepath"

	"github.com/92hackers/code-talks/internal"
)

// NewFile creates a new CodeFile
func NewCodeFile(path string) *internal.CodeFile {
	if filepath.IsAbs(path) == false {
		path, _ = filepath.Abs(path)
	}

	dir, file := filepath.Split(path)

	fileExt := filepath.Ext(file)

	// Get file stats, Follow symbolic links to get the real file stats
	fileInfo, err := os.Stat(path)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	codeFile := &internal.CodeFile{
		FileMetadata: FileMetadata{
			Name:           file,
			Path:           path,
			Directory:      dir,
			FileType:       CODE_FILE,
			LastModifiedAt: uint64(fileInfo.ModTime().Unix()),
		},
		FileContent: FileContent{
			Size:    0,
			Content: "",
		},
		CodeCount:    0,
		CommentCount: 0,
		BlankCount:   0,
		Language:     internal.SupportedLanguages[fileExt],
	}

	// Store the code file
	internal.AllCodeFiles = append(internal.AllCodeFiles, codeFile)

	return codeFile
}

func (f *internal.CodeFile) Analyze() *internal.CodeFile {
	content, err := os.ReadFile(f.Path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	f.Content = string(content)
	f.Size = uint64(len(content))

	return f
}
