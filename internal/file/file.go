/**
Code file
*/

package file

import (
	"os"
	"path/filepath"

	"github.com/92hackers/code-talks/internal"
)

// File types
const (
	CODE_FILE = iota
	CONFIG_FILE
)

type FileMetadata struct {
	// File metadata
	Name           string `json:"name"`
	Path           string `json:"path"`
	Directory      string `json:"directory"`
	FileType       uint8  `json:"file_type"`
	LastModifiedAt uint64 `json:"last_modified_at"`
}

type FileContent struct {
	Size    uint64 `json:"size"`
	Content string `json:"content"`
}

type CodeFile struct {
	FileMetadata
	FileContent

	// Cloc data
	CodeCount    uint32 `json:"code"`
	CommentCount uint32 `json:"comment_count"`
	BlankCount   uint32 `json:"blank_count"`
	TotalLines   uint32 `json:"total_lines"`

	// Code language
	Language string `json:"language"`
}

var AllCodeFiles = []*CodeFile{}

type ConfigFile struct {
	FileMetadata
	FileContent
}

var AllConfigFiles = []*ConfigFile{}

// NewFile creates a new CodeFile
func NewCodeFile(path string) (*CodeFile, error) {
	if filepath.IsAbs(path) == false {
		path, _ = filepath.Abs(path)
	}

	dir, file := filepath.Split(path)

	fileExt := filepath.Ext(file)

	// Get file stats, Follow symbolic links to get the real file stats
	fileInfo, err := os.Stat(path)

	if err != nil {
		return nil, err
	}

	codeFile := &CodeFile{
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
		TotalLines:   0,
		Language:     internal.SupportedLanguages[fileExt].Name,
	}

	// Store the code file
	AllCodeFiles = append(AllCodeFiles, codeFile)

	return codeFile, nil
}

func (f *CodeFile) Analyze() (*CodeFile, error) {
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}

	f.Content = string(content)
	f.Size = uint64(len(content))

	return f, nil
}
