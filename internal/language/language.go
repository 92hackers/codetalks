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

// LanguageAggregateStats represents the aggregate statistics of all programming languages.
type LanguageAggregateStats struct {
	TotalFiles   uint32 `json:"total_files"`
	TotalLines   uint32 `json:"total_lines"`
	TotalCode    uint32 `json:"total_code"`
	TotalComment uint32 `json:"total_comment"`
	TotalBlank   uint32 `json:"total_blank"`
}

// All programming language types detected by codetalks.
var AllLanguages []*Language

// { language-name: Language-index in AllLanguages list }
var AllLanguagesMap = map[string]uint{}

// Aggregate stats for all programming languages.
var AllLanguageAggregateStats = LanguageAggregateStats{}

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

	AllLanguages = append(AllLanguages, language)
	AllLanguagesMap[language.Name] = uint(len(AllLanguages) - 1)

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
	if len(AllLanguages) == 0 {
		return nil
	}
	name := internal.SupportedLanguages[fileExtension].Name
	if _, ok := AllLanguagesMap[name]; !ok {
		return nil
	}
	return AllLanguages[AllLanguagesMap[name]]
}

func AddLanguage(fileExtension string, file *file.CodeFile) *Language {
	language := GetLanguage(fileExtension)
	if language == nil {
		language = NewLanguage(fileExtension)
	}
	language.AddCodeFile(file)
	return language
}

// Aggregate stats
func AggreateStats() {
	for _, language := range AllLanguages {
		AllLanguageAggregateStats.TotalFiles += language.FileCount
		AllLanguageAggregateStats.TotalLines += language.TotalLines
		AllLanguageAggregateStats.TotalCode += language.CodeCount
		AllLanguageAggregateStats.TotalComment += language.CommentCount
		AllLanguageAggregateStats.TotalBlank += language.BlankCount
	}
}

// AnalyzeAllLanguages analyzes all code files and accumulates the data.
func AnalyzeAllLanguages() {
	var wg sync.WaitGroup

	for _, language := range AllLanguages {
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

	// Aggregate
	AggreateStats()

	if internal.IsDebugEnabled {
		fmt.Println("Analyzed all code files")
	}
}
