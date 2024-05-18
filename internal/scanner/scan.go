/**

Scanner.

*/

package scanner

import (
	// "fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/92hackers/code-talks/internal"
	"github.com/92hackers/code-talks/internal/file"
	"github.com/92hackers/code-talks/internal/language"
)

func isVCSDir(path string) bool {
	vcsDirs := []string{".git", ".svn", ".hg", ".bzr", ".cvs"}
	for _, dir := range vcsDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}
	return false
}

func handler(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	leaf := filepath.Base(path)

	// dir
	if d.IsDir() {
		// Skip VCS directories
		if isVCSDir(leaf) {
			return fs.SkipDir
		}

		// TODO: Skip specified-ignored directories

		return nil
	}

	// Skip unsupported file extensions
	// TODO: maybe a config file
	fileExt := filepath.Ext(leaf)
	if internal.SupportedLanguages[fileExt] == nil {
		return nil
	}

	// Create a new code file, skip if error
	codeFile, err := file.NewCodeFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Add the code file to the language
	language.AddLanguage(fileExt, codeFile)

	return nil
}

func Scan(dir string) {
	// Scan directory
	err := filepath.WalkDir(dir, handler)
	if err != nil {
		log.Fatal(err)
	}
}
