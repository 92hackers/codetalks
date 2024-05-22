/**
Internal
*/

package internal

// LanguageDefinition represents a programming language definition.
type LanguageDefinition struct {
	Name          string     `json:"name"`
	LineComments  []string   `json:"line_comment"`
	BlockComments [][]string `json:"block_comment"`
}

func newLangDef(name string, lineComments []string, blockComments [][]string) *LanguageDefinition {
	return &LanguageDefinition{
		Name:          name,
		LineComments:  lineComments,
		BlockComments: blockComments,
	}
}

// Supported programming languages, map as file extension -> language name.
var SupportedLanguages = map[string]*LanguageDefinition{
	".s":   newLangDef("Assembly", []string{"//"}, [][]string{{"/*", "*/"}}),
	".c":   newLangDef("C", []string{"//"}, [][]string{{"/*", "*/"}}),
	".h":   newLangDef("C Header", []string{"//"}, [][]string{{"/*", "*/"}}),
	".hh":  newLangDef("C++ Header", []string{"//"}, [][]string{{"/*", "*/"}}),
	".hpp": newLangDef("C++ Header", []string{"//"}, [][]string{{"/*", "*/"}}),
	".cpp": newLangDef("C++", []string{"//"}, [][]string{{"/*", "*/"}}),
	".cs":  newLangDef("C#", []string{"//"}, [][]string{{"/*", "*/"}}),
	".rs":  newLangDef("Rust", []string{"//"}, [][]string{{"/*", "*/"}}),
	".lua": newLangDef("Lua", []string{"--"}, [][]string{{"--[[", "]]"}}),

	".js":  newLangDef("JavaScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".mjs": newLangDef("JavaScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".cjs": newLangDef("JavaScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".jsx": newLangDef("JavaScript", []string{"//"}, [][]string{{"/*", "*/"}}),

	".ts":  newLangDef("TypeScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".mts": newLangDef("TypeScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".cts": newLangDef("TypeScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".tsx": newLangDef("TypeScript", []string{"//"}, [][]string{{"/*", "*/"}}),

	".java":  newLangDef("Java", []string{"//"}, [][]string{{"/*", "*/"}}),
	".php":   newLangDef("PHP", []string{"//"}, [][]string{{"/*", "*/"}}),
	".py":    newLangDef("Python", []string{"#"}, [][]string{{"\"\"\"", "\"\"\""}}),
	".pl":    newLangDef("Perl", []string{"#"}, [][]string{{":=", ":=cut"}}),
	".rb":    newLangDef("Ruby", []string{"#"}, [][]string{{":=begin", ":=end"}}),
	".swift": newLangDef("Swift", []string{"//"}, [][]string{{"/*", "*/"}}),
	".go":    newLangDef("Go", []string{"//"}, [][]string{{"/*", "*/"}}),
	".kt":    newLangDef("Kotlin", []string{"//"}, [][]string{{"/*", "*/"}}),
	".scala": newLangDef("Scala", []string{"//"}, [][]string{{"/*", "*/"}}),
	".r":     newLangDef("R", []string{"#"}, [][]string{{"/*", "*/"}}),

	".sh":   newLangDef("Shell", []string{"#"}, [][]string{{"", ""}}),
	".zsh":  newLangDef("Shell", []string{"#"}, [][]string{{"", ""}}),
	".bash": newLangDef("Shell", []string{"#"}, [][]string{{"", ""}}),

	".html": newLangDef("HTML", []string{"<!--", "//"}, [][]string{{"<!--", "-->"}}),
	".xml":  newLangDef("XML", []string{"<!--"}, [][]string{{"<!--", "-->"}}),
	".css":  newLangDef("CSS", []string{"//"}, [][]string{{"/*", "*/"}}),

	".json": newLangDef("JSON", []string{}, [][]string{{"", ""}}),
	".yaml": newLangDef("YAML", []string{"#"}, [][]string{{"", ""}}),
	".toml": newLangDef("TOML", []string{"#"}, [][]string{{"", ""}}),

	".md":  newLangDef("Markdown", []string{}, [][]string{{"", ""}}),
	".txt": newLangDef("Plain Text", []string{}, [][]string{{"", ""}}),
}

// Config files, map as file name -> file label.
// These files are commonly used in programming projects.
var ConfigFiles = map[string]string{
	"makefile": "Makefile",
	"rakefile": "Rakefile",
	"gemfile":  "Gemfile",

	// Version control
	".gitignore":  ".gitignore",
	".gitmodules": ".gitmodules",
	".gitconfig":  ".gitconfig",

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

// Debug mode
var IsDebugEnabled = false
