/**

Process single file.

*/

package file

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

	// Code language
	Language string `json:"language"`
}

var AllCodeFiles []*CodeFile

type ConfigFile struct {
	FileMetadata
	FileContent
}

var AllConfigFiles []*ConfigFile
