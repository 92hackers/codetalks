/**

Supported programming languages.

*/

package language

import (
  "regexp"
)

type Language struct {
  // Language metadata
  Name string `json:"name"`
  Lable string `json:"label"`

	// Cloc data
	CodeCount     int32 `json:"code"`
	CommentCount int32 `json:"comment_count"`
	BlankCount   int32 `json:"blank_count"`

  FileExtension string `json:"file_extension"`
  FilesCount int32 `json:"files_count"`
}

// Supported languages, map as file extension -> language label.
var SupportedLanguages = map[string]string{
  "c": "C",
  "cpp": "C++",
  "cs": "C#",
  "java": "Java",
  "js": "JavaScript",
  "php": "PHP",
  "py": "Python",
  "rb": "Ruby",
  "rs": "Rust",
  "swift": "Swift",
  "go": "Go",
  "kt": "Kotlin",
  "ts": "TypeScript",
  "scala": "Scala",
  "r": "R",
  "sh": "Shell",
  "pl": "Perl",
  "lua": "Lua",
  "html": "HTML",
  "css": "CSS",
  "xml": "XML",
  "json": "JSON",
  "yaml": "YAML",
  "toml": "TOML",
  "md": "Markdown",
  "txt": "Text",
  "makefile": "Makefile",
}

