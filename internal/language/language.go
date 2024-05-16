/**

Supported programming languages.

*/

package language

import (
	// "regexp"
	"github.com/92hackers/code-talks/internal/file"
)

func NewLanguage(label, fileExtension string) *Language {
	language := &Language{
		Name:          name,
		Lable:         label,
		FileExtension: fileExtension,
	}

	AllLanguages = append(AllLanguages, language)

	return language
}
