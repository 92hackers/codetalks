/**
Code file
*/

package file

import (
	"bufio"
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/92hackers/codetalks/internal"
)

// File types
type FileType uint8

const (
	CODE_FILE FileType = iota // Code file
	CONFIG_FILE
)

// Struct memory alignment consideration
type FileMetadata struct {
	// File metadata
	LastModifiedAt uint64   `json:"last_modified_at"`
	FileType       FileType `json:"file_type"`
	Name           string   `json:"name"`
	Path           string   `json:"path"`
	Directory      string   `json:"directory"`
	FileExtension  string   `json:"file_extension"`
}

type FileContent struct {
	Size    uint64 `json:"size"`
	Content string `json:"content"`
}

type CodeFile struct {
	// Cloc data
	CodeCount    uint32 `json:"code"`
	CommentCount uint32 `json:"comment_count"`
	BlankCount   uint32 `json:"blank_count"`
	TotalLines   uint32 `json:"total_lines"`

	FileMetadata
	FileContent

	// Code language
	Language string `json:"language"`
}

// A mutex to protect the AllCodeFiles.
var AllCodeFilesLock sync.Mutex
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
		Language: internal.SupportedLanguages[fileExt].Name,
	}

	// Store the code file
	addFile(codeFile)

	return codeFile, nil
}

func addFile(codeFile *CodeFile) {
	AllCodeFilesLock.Lock()
	defer AllCodeFilesLock.Unlock()
	AllCodeFiles = append(AllCodeFiles, codeFile)
}

type stateFlag struct {
	isInBlockComment bool
	isFirstLine      bool
}

func (f *CodeFile) Analyze(ctx context.Context) (*CodeFile, error) {
	// Open file
	fd, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	scanState := stateFlag{}
	// Init scanner buffer as 64KB, max token size as 10MB
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			log.Println("Analyze timeout, skipping file: ", f.Path)
			return nil, nil

		default:
			line := scanner.Text()
			f.scanLine(line, &scanState)
		}
	}

	return f, nil
}

func (f *CodeFile) scanLine(line string, state *stateFlag) {
	f.TotalLines++

	line = strings.TrimSpace(line)

	f.AddFileContent(line)

	langDefinition := internal.SupportedLanguages[f.FileExtension]

	// shebang, treated as code
	if state.isFirstLine == false && strings.HasPrefix(line, "#!") {
		f.CodeCount++
		state.isFirstLine = true
		return
	}

	if state.isInBlockComment == true {
		f.CommentCount++
		// Check if the block comment ends in this line
		for _, comment := range langDefinition.BlockComments {
			end := comment[1]
			if len(end) > 0 && strings.Contains(line, end) {
				state.isInBlockComment = false
				break
			}
		}
		return
	}

	// Blank line
	if line == "" {
		f.BlankCount++
		return
	}

	// Single line comment
	for _, comment := range langDefinition.LineComments {
		if len(comment) > 0 && strings.HasPrefix(line, comment) {
			f.CommentCount++
			return
		}
	}

	// Note: currently, nested block comments are not parsed..
	// Block comment begin
	for _, comment := range langDefinition.BlockComments {
		begin, end := comment[0], comment[1]
		if len(begin) > 0 && strings.HasPrefix(line, begin) {
			// 1. Begining of a new block comment, if end is not in the same line.
			if !strings.Contains(line, end) {
				state.isInBlockComment = true
			}
			f.CommentCount++
			return
		}
	}

	// Code line
	f.CodeCount++
}

func (f *CodeFile) AddFileContent(content string) {
	// f.Content += content
	f.Size += uint64(len(content))
}
