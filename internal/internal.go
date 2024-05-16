/**
Internal
*/

package internal

// Language represents a programming language.
type Language struct {
	// Language metadata
	FileExtension string `json:"file_extension"`
	Lable         string `json:"label"`

	// Cloc data
	CodeCount    uint32 `json:"code"`
	CommentCount uint32 `json:"comment_count"`
	BlankCount   uint32 `json:"blank_count"`
	TotalLines   uint32 `json:"total_lines"`

	FileCount uint32 `json:"file_count"`
	CodeFiles []*CodeFile
}

// All programming language types detected by codetalks.
var AllLanguagesMap map[string]*Language

// Supported programming languages, map as file extension -> language label.
const SupportedLanguages = map[string]string{
	".c":     "C",
	".cpp":   "C++",
	".cs":    "C#",
	".java":  "Java",
	".js":    "JavaScript",
	".php":   "PHP",
	".py":    "Python",
	".rb":    "Ruby",
	".rs":    "Rust",
	".swift": "Swift",
	".go":    "Go",
	".kt":    "Kotlin",
	".ts":    "TypeScript",
	".scala": "Scala",
	".r":     "R",
	".sh":    "Shell",
	".pl":    "Perl",
	".lua":   "Lua",
	".html":  "HTML",
	".css":   "CSS",
	".xml":   "XML",
	".json":  "JSON",
	".yaml":  "YAML",
	".toml":  "TOML",
	".md":    "Markdown",
	".txt":   "Text",
}

// Config files, map as file name -> file label.
// These files are commonly used in programming projects.
const ConfigFiles = map[string]string{
	"makefile": "Makefile",
	"rakefile": "Rakefile",
	"gemfile":  "Gemfile",

	// Docker
	"dockerfile":         "Dockerfile",
	"docker-compose.yml": "Docker Compose file",

	// JavaScript
	"package.json":  "package.json",
	"yarn.lock":     "yarn.lock",
	"tsconfig.json": "tsconfig.json",

	// Go
	"go.mod": "go.mod",
	"go.sum": "go.sum",
}

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
