/**

Scanner.

*/

package scanner

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/92hackers/codetalks/internal"
	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
	"github.com/92hackers/codetalks/internal/utils"
	"github.com/92hackers/codetalks/internal/view_mode"
)

var (
	uniqueDirSet *utils.Set
	matchRegex   []*regexp.Regexp
	ignoreRegex  []*regexp.Regexp
)

func init() {
	// Initialize the unique directory set
	uniqueDirSet = utils.NewSet()
}

func Config(
	matchRegexStr string,
	ignoreRegexStr string,
) {
	// Match regexs
	matchRegexStr = strings.TrimSpace(matchRegexStr)
	if len(matchRegexStr) > 0 {
		for _, regexStr := range strings.Split(matchRegexStr, " ") {
			matchRegex = append(matchRegex, regexp.MustCompile(regexStr))
		}
	}

	// Ignore regexs
	ignoreRegexStr = strings.TrimSpace(ignoreRegexStr)
	if len(ignoreRegexStr) > 0 {
		for _, regexStr := range strings.Split(ignoreRegexStr, " ") {
			ignoreRegex = append(ignoreRegex, regexp.MustCompile(regexStr))
		}
	}
}

func isSpecifiedDepthDirs(path string, depth int) bool {
	segments := make([]string, 0, 10)
	for {
		dir, file := filepath.Split(path)
		trimedFile := strings.TrimSpace(file)
		if trimedFile != "" {
			segments = append(segments, trimedFile)
		}
		if dir == string(os.PathSeparator) || dir == "" {
			break
		}
		if strings.HasSuffix(dir, string(os.PathSeparator)) {
			dir = strings.TrimSuffix(dir, string(os.PathSeparator))
		}
		path = dir
	}
	return len(segments) == depth
}

func handler(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	leaf := filepath.Base(path)

	// Cut the root directory from the scanned path.
	cutRootDirPath := strings.TrimPrefix(path, utils.CurrentRootDir)

	// dir
	if d.IsDir() {
		// Store the directory if viewMode is set to directory
		if internal.GlobalOpts.ViewMode == view_mode.ViewModeDirs {
			// Check Depth and store the directory
			// TODO: support multiple depth
			if isSpecifiedDepthDirs(cutRootDirPath, 1) {
				view_mode.SubDirs = append(view_mode.SubDirs, cutRootDirPath)
			}
		}

		return nil
	}

	// TODO: handle config file

	// Match regex filter
	{
		isMatched := false

		for _, re := range matchRegex {
			if re.MatchString(cutRootDirPath) {
				isMatched = true
				// Log
				if internal.GlobalOpts.IsDebugEnabled || internal.GlobalOpts.IsShowMatched {
					log.Println("File matched:", path)
				}
				break
			}
			if internal.GlobalOpts.IsDebugEnabled {
				log.Println("Not matched:", path, "with regexp:", re.String())
			}
		}
		if len(matchRegex) > 0 && !isMatched {
			return nil
		}

		// Matched by matchRegex
		// Custom match regular expression has over precedence over gitignore patterns

		if !isMatched {
			// Check if the file is ignored by gitignore
			if utils.ShouldIgnoreFile(cutRootDirPath) {
				if internal.GlobalOpts.IsDebugEnabled {
					log.Println("File ignored by .gitignore rules:", path)
				}
				return nil
			}
		}
	}

	// Ignore regex filter
	for _, re := range ignoreRegex {
		if re.MatchString(cutRootDirPath) {
			if internal.GlobalOpts.IsDebugEnabled || internal.GlobalOpts.IsShowIgnored {
				log.Println("File ignored:", path, "with regexp:", re.String())
			}
			return nil
		}
	}

	// Skip unsupported file extensions
	fileExt := filepath.Ext(leaf)
	if internal.SupportedLanguages[fileExt] == nil {
		if internal.GlobalOpts.IsDebugEnabled {
			utils.ErrorMsg("Unsupported file type: %s", path)
		}
		return nil
	}

	// Duplicate directory check
	if uniqueDirSet.Contains(path) {
		return nil
	}

	// debug
	if internal.GlobalOpts.IsDebugEnabled {
		log.Println("Add new file:", path)
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
		if err := utils.WalkDir(dir, handler); err != nil {
			log.Fatal(err)
		}
	}
}
