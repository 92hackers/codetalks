/**
Code Talks
*/

package codeTalks

// Supported programming languages, map as file extension -> language label.
var SupportedLanguages = map[string]string{
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
var ConfigFiles = map[string]string{
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
