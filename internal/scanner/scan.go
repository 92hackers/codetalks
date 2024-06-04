/**

Scanner.

*/

package scanner

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/92hackers/codetalks/internal"
	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
	"github.com/92hackers/codetalks/internal/utils"
)

var (
	uniqueDirSet   *utils.Set
	matchRegex     []*regexp.Regexp
	ignoreRegex    []*regexp.Regexp
	currentRootDir string
	vcsDirs        *utils.Set
)

func init() {
	// Initialize the unique directory set
	uniqueDirSet = utils.NewSet()
	vcsDirs = utils.NewSet()
	{
		vcsDirs.Add(".git")
		vcsDirs.Add(".svn")
		vcsDirs.Add(".hg")
		vcsDirs.Add(".bzr")
		vcsDirs.Add(".cvs")
	}
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

func isVCSDir(path string) bool {
	return vcsDirs.Contains(path)
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

	// Cut the root directory from the scanned path.
	cutRootDirPath := strings.TrimPrefix(path, currentRootDir)
	// Match regex filter
	{
		isMatched := false
		for _, re := range matchRegex {
			if re.MatchString(cutRootDirPath) {
				isMatched = true
				// Log
				if internal.GlobalOpts.IsDebugEnabled || internal.GlobalOpts.IsShowMatched {
					fmt.Println("File matched:", path)
				}
				break
			}
			if internal.GlobalOpts.IsDebugEnabled {
				fmt.Println("Not matched:", path, "with regexp:", re.String())
			}
		}
		if len(matchRegex) > 0 && !isMatched {
			return nil
		}
	}
	// Ignore regex filter
	for _, re := range ignoreRegex {
		if re.MatchString(cutRootDirPath) {
			if internal.GlobalOpts.IsDebugEnabled || internal.GlobalOpts.IsShowIgnored {
				fmt.Println("File ignored:", path, "with regexp:", re.String())
			}
			return nil
		}
	}

	// Skip unsupported file extensions
	fileExt := filepath.Ext(leaf)
	if internal.SupportedLanguages[fileExt] == nil {
		if internal.GlobalOpts.IsDebugEnabled {
			fmt.Println("❌ Unsupported file type:", path)
		}
		return nil
	}

	// Duplicate directory check
	if uniqueDirSet.Contains(path) {
		return nil
	}

	// debug
	if internal.GlobalOpts.IsDebugEnabled {
		fmt.Println("Add new file:", path)
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
		currentRootDir = dir
		// Scan directory
		err := filepath.WalkDir(dir, handler)
		if err != nil {
			log.Fatal(err)
		}
	}
}
