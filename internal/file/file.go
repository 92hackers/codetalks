/**

Process single file.

*/

package file

type CodeTalksFile struct {
  // Cloc data
  Code int32 `json:"code"`
  Comments int32 `json:"comments"`
  Blanks int32 `json:"blanks"`

  // File metadata
  Name string `json:"name"`
  Path string `json:"path"`
  Language string `json:"language"`
}

func (f *CodeTalksFile) GetCode() int32 {
  return f.Code
}
