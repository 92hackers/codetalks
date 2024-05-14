/**

Process single file.

*/

package file

type CodeTalksFile struct {
	// Cloc data
	CodeCount     int32 `json:"code"`
	CommentCount int32 `json:"comment_count"`
	BlankCount   int32 `json:"blank_count"`

	// File metadata
	Name     string `json:"name"`
	Path     string `json:"path"`
	Language string `json:"language"`
}

type CodeTalksFiles []CodeTalksFile

func (f *CodeTalksFile) GetCode() int32 {
	return f.Code
}

