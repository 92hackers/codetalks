/**

Supported programming languages.

*/

package language

import (
	"github.com/92hackers/code-talks/internal"
)

type Language internal.Language
type CodeFile internal.CodeFile

func NewLanguage(fileExtension string) *Language {
	language := &Language{
		Name:          internal.SupportedLanguages[fileExtension].Name,
		FileExtension: fileExtension,
    CodeCount:     0,
    CommentCount:  0,
    BlankCount:    0,
    TotalLines:    0,
    FileCount:     0,
    CodeFiles:     []*CodeFile{},
	}

  internal.AllLanguagesMap[language.Name] = language

	return language
}

func (l *Language) AddCodeFile(file *CodeFile) *Language {
  l.CodeCount += file.CodeCount
  l.CommentCount += file.CommentCount
  l.BlankCount += file.BlankCount
  l.TotalLines += file.TotalLines
  l.FileCount += 1
  l.CodeFiles = append(l.CodeFiles, file)
  return l
}

func GetLanguage(fileExtension string) *Language {
  name := internal.SupportedLanguages[fileExtension].Name
  return internal.AllLanguagesMap[name]
}

func AddLanguage(fileExtension string, file *CodeFile) *Language {
  language := GetLanguage(fileExtension)
  if language == nil {
    language = NewLanguage(fileExtension)
  }
  language.AddCodeFile(file)
  return language
}
