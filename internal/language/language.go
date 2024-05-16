/**

Supported programming languages.

*/

package language

import (
	// "regexp"
	"github.com/92hackers/code-talks/internal/file"
)

type Language struct {
	// Language metadata
	Name  string `json:"name"`
	Lable string `json:"label"`

	// Cloc data
	CodeCount    uint32 `json:"code"`
	CommentCount uint32 `json:"comment_count"`
	BlankCount   uint32 `json:"blank_count"`

	FileExtension string `json:"file_extension"`
	FileCount     uint32 `json:"file_count"`
	CodeFiles     []*file.CodeFile
}

var AllLanguages []*Language
