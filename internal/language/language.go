/**

Supported programming languages.

*/

package language

import (
	"fmt"
	"log"
	"sync"

	"github.com/92hackers/code-talks/internal"
	"github.com/92hackers/code-talks/internal/file"
)

// Language represents a programming language.
type Language struct {
	// Language metadata
	Name          string `json:"name"`
	FileExtension string `json:"file_extension"`

	// Cloc data
	CodeCount    uint32 `json:"code"`
	CommentCount uint32 `json:"comment_count"`
	BlankCount   uint32 `json:"blank_count"`
	TotalLines   uint32 `json:"total_lines"`

	FileCount uint32 `json:"file_count"`
	CodeFiles []*file.CodeFile
}

// All programming language types detected by codetalks.
// { language-name: Language }
var AllLanguagesMap = map[string]*Language{}

func NewLanguage(fileExtension string) *Language {
	language := &Language{
		Name:          internal.SupportedLanguages[fileExtension].Name,
		FileExtension: fileExtension,
		CodeCount:     0,
		CommentCount:  0,
		BlankCount:    0,
		TotalLines:    0,
		FileCount:     0,
	}

	AllLanguagesMap[language.Name] = language

	return language
}

func (l *Language) AddCodeFile(file *file.CodeFile) *Language {
	l.FileCount += 1
	l.CodeFiles = append(l.CodeFiles, file)
	return l
}

func (l *Language) CountCodeFileStats(file *file.CodeFile) *Language {
	l.CodeCount += file.CodeCount
	l.CommentCount += file.CommentCount
	l.BlankCount += file.BlankCount
	l.TotalLines += file.TotalLines
	return l
}

func GetLanguage(fileExtension string) *Language {
	name := internal.SupportedLanguages[fileExtension].Name
	return AllLanguagesMap[name]
}

func AddLanguage(fileExtension string, file *file.CodeFile) *Language {
	language := GetLanguage(fileExtension)
	if language == nil {
		language = NewLanguage(fileExtension)
	}
	language.AddCodeFile(file)
	return language
}

func AnalyzeAllLanguagesSlow() {
	for _, language := range AllLanguagesMap {
		for _, codeFile := range language.CodeFiles {
			f, err := codeFile.Analyze()
			if err != nil {
				continue
			}
			language.CountCodeFileStats(f)
		}
	}

	fmt.Println("Analyzed all code files")
}

// AnalyzeAllLanguages analyzes all code files and accumulates the data.
func AnalyzeAllLanguages() {
	var wg sync.WaitGroup

	for _, language := range AllLanguagesMap {
		for _, codeFile := range language.CodeFiles {
			wg.Add(1)
			go func(codeFile *file.CodeFile) {
				// TODO: Add error handling, handle timeout, and cancelation, maybe by <-done channell.
				f, err := codeFile.Analyze()
				if err != nil {
					log.Println(err) // Log error and continue.
					return
				}
				language.CountCodeFileStats(f)
				wg.Done()
			}(codeFile)
		}
	}

	wg.Wait()
	fmt.Println("Analyzed all code files")
}
