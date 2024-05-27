/**

Scanner.

*/

package scanner

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/92hackers/codetalks/internal"
	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
	"github.com/92hackers/codetalks/internal/utils"
)

var uniqueDirSet *utils.Set

func init() {
	// Initialize the unique directory set
  uniqueDirSet = utils.NewSet()
}

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

	// TODO: handle config file

	// Skip unsupported file extensions
	fileExt := filepath.Ext(leaf)
	if internal.SupportedLanguages[fileExt] == nil {
		if internal.IsDebugEnabled {
			fmt.Println("❌ Unsupported file type", path)
		}
		return nil
	}

	// Duplicate directory check
	if uniqueDirSet.Contains(path) {
		return nil
	}

	// debug
	if internal.IsDebugEnabled {
		fmt.Println("Add new file: ", path)
	}

	// Create a new code file, skip if error
	codeFile, err := file.NewCodeFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Add the code file to the language
	language.AddLanguage(fileExt, codeFile)

	// Add the directory to the unique directory set
	uniqueDirSet.Add(path)

	return nil
}

func Scan(rootDirs []string) {
	for _, dir := range rootDirs {
		// Scan directory
		err := filepath.WalkDir(dir, handler)
		if err != nil {
			log.Fatal(err)
		}
	}
}
