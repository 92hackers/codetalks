/**

Supported programming languages.

*/

package language

import (
	"fmt"
	"log"
	"sync"

	"github.com/92hackers/codetalks/internal"
	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/utils"
)

// Language represents a programming language.
type Language struct {
	// Cloc data
	CodeCount    uint32 `json:"code"`
	CommentCount uint32 `json:"comment_count"`
	BlankCount   uint32 `json:"blank_count"`
	TotalLines   uint32 `json:"total_lines"`

	FileCount uint32 `json:"file_count"`
	CodeFiles []*file.CodeFile

	// Language metadata
	Name          string `json:"name"`
	FileExtension string `json:"file_extension"`
}

// LanguageAggregateStats represents the aggregate statistics of all programming languages.
type LanguageAggregateStats struct {
	TotalFiles   uint32 `json:"total_files"`
	TotalLines   uint32 `json:"total_lines"`
	TotalCode    uint32 `json:"total_code"`
	TotalComment uint32 `json:"total_comment"`
	TotalBlank   uint32 `json:"total_blank"`
}

// A mutex to protect the AllLanguages and AllLanguagesMap.
var AllLangsLock sync.Mutex

// All programming language types detected by codetalks.
var AllLanguages []*Language

// { language-name: Language-index in AllLanguages list }
var AllLanguagesMap = map[string]uint{}

// Aggregate stats for all programming languages.
var AllLanguageAggregateStats = LanguageAggregateStats{}

func newLanguage(fileExtension string) *Language {
	language := &Language{
		Name:          internal.SupportedLanguages[fileExtension].Name,
		FileExtension: fileExtension,
	}

	AllLanguages = append(AllLanguages, language)
	AllLanguagesMap[language.Name] = uint(len(AllLanguages) - 1)

	return language
}

func (l *Language) addCodeFile(file *file.CodeFile) *Language {
	l.FileCount += 1
	l.CodeFiles = append(l.CodeFiles, file)
	return l
}

// CountCodeFileStats counts the code, comment, and blank lines in a code file.
// It's not thread-safe.
func (l *Language) CountCodeFileStats(file *file.CodeFile) *Language {
	if file == nil {
		return l
	}
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
	langDef, ok := internal.SupportedLanguages[fileExtension]
	if !ok {
		return nil
	}
	name := langDef.Name
	if _, ok := AllLanguagesMap[name]; !ok {
		return nil
	}
	return AllLanguages[AllLanguagesMap[name]]
}

func AddLanguage(fileExtension string, file *file.CodeFile) *Language {
	AllLangsLock.Lock()
	defer AllLangsLock.Unlock()

	language := GetLanguage(fileExtension)
	if language == nil {
		language = newLanguage(fileExtension)
	}
	language.addCodeFile(file)
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
		// A channel used to receive the analyzed files
		ch := make(chan *file.CodeFile)

		for _, codeFile := range language.CodeFiles {
			// Every file has 1s to analyze itself.
			// ctx, cancelCtx := utils.WithTimeoutCtxMilliSeconds(300)
			ctx, cancelCtx := utils.WithTimeoutCtxSeconds(1)
			defer cancelCtx()

			wg.Add(1)
			go func(codeFile *file.CodeFile, ch chan<- *file.CodeFile) {
				defer wg.Done()
				f, err := codeFile.Analyze(ctx)
				if err != nil {
					log.Println(err)
					ch <- nil
					return
				}
				ch <- f
			}(codeFile, ch)
		}

		// Aggregate stats for the language
		wg.Add(1)
		go func(language *Language, ch chan *file.CodeFile) {
			defer close(ch)
			for i := 0; i < len(language.CodeFiles); i++ {
				language.CountCodeFileStats(<-ch)
			}
			wg.Done()
		}(language, ch)
	}

	wg.Wait()

	// Aggregate
	AggreateStats()

	if internal.IsDebugEnabled {
		fmt.Println("All code files analyzed!")
	}
}
