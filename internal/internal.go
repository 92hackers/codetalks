/**
Internal
*/

package internal

// LanguageDefinition represents a programming language definition.
type LanguageDefinition struct {
	Name         string     `json:"name"`
	lineComment  []string   `json:"line_comment"`
	blockComment [][]string `json:"block_comment"`
}

func newLangDef(name string, lineComment []string, blockComment [][]string) *LanguageDefinition {
	return &LanguageDefinition{
		Name:         name,
		lineComment:  lineComment,
		blockComment: blockComment,
	}
}

// Supported programming languages, map as file extension -> language name.
var SupportedLanguages = map[string]*LanguageDefinition{
	".c":     newLangDef("C", []string{"//"}, [][]string{{"/*", "*/"}}),
	".h":     newLangDef("C Header", []string{"//"}, [][]string{{"/*", "*/"}}),
	".hh":    newLangDef("C++ Header", []string{"//"}, [][]string{{"/*", "*/"}}),
	".hpp":   newLangDef("C++ Header", []string{"//"}, [][]string{{"/*", "*/"}}),
	".cpp":   newLangDef("C++", []string{"//"}, [][]string{{"/*", "*/"}}),
	".cs":    newLangDef("C#", []string{"//"}, [][]string{{"/*", "*/"}}),
	".java":  newLangDef("Java", []string{"//"}, [][]string{{"/*", "*/"}}),
	".js":    newLangDef("JavaScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".mjs":   newLangDef("JavaScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".ts":    newLangDef("TypeScript", []string{"//"}, [][]string{{"/*", "*/"}}),
	".php":   newLangDef("PHP", []string{"//"}, [][]string{{"/*", "*/"}}),
	".py":    newLangDef("Python", []string{"#"}, [][]string{{"\"\"\"", "\"\"\""}}),
	".rb":    newLangDef("Ruby", []string{"#"}, [][]string{{":=begin", ":=end"}}),
	".rs":    newLangDef("Rust", []string{"//"}, [][]string{{"/*", "*/"}}),
	".swift": newLangDef("Swift", []string{"//"}, [][]string{{"/*", "*/"}}),
	".go":    newLangDef("Go", []string{"//"}, [][]string{{"/*", "*/"}}),
	".kt":    newLangDef("Kotlin", []string{"//"}, [][]string{{"/*", "*/"}}),
	".scala": newLangDef("Scala", []string{"//"}, [][]string{{"/*", "*/"}}),
	".r":     newLangDef("R", []string{"#"}, [][]string{{"/*", "*/"}}),
	".sh":    newLangDef("Shell", []string{"#"}, [][]string{{"", ""}}),
	".pl":    newLangDef("Perl", []string{"#"}, [][]string{{":=", ":=cut"}}),
	".lua":   newLangDef("Lua", []string{"--"}, [][]string{{"--[[", "]]"}}),
	".html":  newLangDef("HTML", []string{"<!--", "//"}, [][]string{{"<!--", "-->"}}),
	".css":   newLangDef("CSS", []string{"//"}, [][]string{{"/*", "*/"}}),
	".xml":   newLangDef("XML", []string{"<!--"}, [][]string{{"<!--", "-->"}}),
	".json":  newLangDef("JSON", []string{}, [][]string{{"", ""}}),
	".yaml":  newLangDef("YAML", []string{"#"}, [][]string{{"", ""}}),
	".toml":  newLangDef("TOML", []string{"#"}, [][]string{{"", ""}}),
	".md":    newLangDef("Markdown", []string{}, [][]string{{"", ""}}),
	".txt":   newLangDef("Plain Text", []string{}, [][]string{{"", ""}}),
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
