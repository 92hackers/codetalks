/**

CodeTalks command

*/

package main

import (
	"fmt"
	"github.com/92hackers/code-talks/internal/file"
	// "os"
)

// RootDir is the root directory of the codebase
var RootDir string

func main() {
	file := &file.CodeFile{
		FileMetadata: file.FileMetadata{
			Name:           "main.go",
			Path:           "go.mod",
			Directory:      "internal",
			FileType:       file.CODE_FILE,
			LastModifiedAt: 100,
		},
		FileContent: file.FileContent{
			Size:    0,
			Content: "",
		},
		CodeCount:    10,
		CommentCount: 10,
		BlankCount:   10,
		Language:     "Go",
	}

	fmt.Println(file.Language)
	fmt.Println(file.LastModifiedAt)

	file.Analyze()

	fmt.Println(file.Content)
	fmt.Println(file.Size)
}
