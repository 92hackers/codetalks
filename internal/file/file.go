/**
Code file
*/

package file

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

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
	FileExtension  string `json:"file_extension"`
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
			FileExtension:  fileExt,
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
	isInBlockComment := false
	isFirstLine := false

	// Open file
	fd, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	// Init scanner buffer as 64KB, max token size as 10MB
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)

	langDefinition := internal.SupportedLanguages[f.FileExtension]

scannerFor:
	for scanner.Scan() {
		line := scanner.Text()
		f.TotalLines++

		line = strings.TrimSpace(line)

		f.AddFileContent(line)

		// shebang, treated as code
		if isFirstLine == false && strings.HasPrefix(line, "#!") {
			f.CodeCount++
			isFirstLine = true
			continue
		}

		if isInBlockComment == true {
			f.CommentCount++
			// Check if the block comment ends in this line
			for _, comment := range langDefinition.BlockComments {
				end := comment[1]
				if len(end) > 0 && strings.Contains(line, end) {
					isInBlockComment = false
					break
				}
			}
			continue
		}

		// Blank line
		if line == "" {
			f.BlankCount++
			continue
		}

		// Single line comment
		for _, comment := range langDefinition.LineComments {
			if len(comment) > 0 && strings.HasPrefix(line, comment) {
				f.CommentCount++
				continue scannerFor
			}
		}

		// Currently, nested block comments are not parsed..

		// Block comment begin
		for _, comment := range langDefinition.BlockComments {
			begin, end := comment[0], comment[1]
			if len(begin) > 0 && strings.HasPrefix(line, begin) {
				// 1. Begining of a new block comment, if end is not in the same line.
				if !strings.Contains(line, end) {
					isInBlockComment = true
				}
				f.CommentCount++
				continue scannerFor
			}
		}

		// Code line
		f.CodeCount++
	}

	return f, nil
}

func (f *CodeFile) AddFileContent(content string) {
	// f.Content += content
	f.Size += uint64(len(content))
}
