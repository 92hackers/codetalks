/**

Walk a directory tree recursively.

A concurrent version of filepath.WalkDir.

Features:

0. Concurrent walk.
1. Respect .gitignore rules.
2. Skip VCS directories.

*/

package utils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/92hackers/codetalks/internal/git"
)

var (
	VCSDirs        *Set
	GitIgnore      *ignore.GitIgnore
	CurrentRootDir string
)

func init() {
	VCSDirs = NewSet()
	dirs := []string{".git", ".svn", ".hg", ".bzr", ".cvs"}
	for _, dir := range dirs {
		VCSDirs.Add(dir)
	}
}

func isVCSDir(path string) bool {
	return VCSDirs.Contains(path)
}

// WalkDir walks the file tree rooted at root, calling fn for each file or
// directory in the tree, including root.
//
// All errors that arise visiting files and directories are filtered by fn:
// see the [fs.WalkDirFunc] documentation for details.
//
// The files are walked in lexical order, which makes the output deterministic
// but requires WalkDir to read an entire directory into memory before proceeding
// to walk that directory.
//
// WalkDir does not follow symbolic links.
//
// WalkDir calls fn with paths that use the separator character appropriate
// for the operating system. This is unlike [io/fs.WalkDir], which always
// uses slash separated paths.
func WalkDir(root string, fn fs.WalkDirFunc) error {
	info, err := os.Lstat(root)
	if err != nil {
		return err
	}

	CurrentRootDir = root

	// Respect .gitignore rules.
	// If parse failed, continue without .gitignore rules.
	ignoreFile := filepath.Join(root, ".gitignore")
	GitIgnore, _ = ignore.CompileIgnoreFile(ignoreFile)

	var wg sync.WaitGroup

	wg.Add(1)
	go walkDir(root, fs.FileInfoToDirEntry(info), fn, &wg)

	wg.Wait()

	return nil
}

func shouldIgnoreDir(path string, d fs.DirEntry) bool {
	if !d.IsDir() {
		return false
	}
	leaf := filepath.Base(path)
	// Cut the root directory from the scanned path.
	cutRootDirPath := strings.TrimPrefix(path, CurrentRootDir)
	// 1. Skip VCS directories
	// 2. Skip directories that are ignored by gitignore
	//
	// Custom match regular expression has over precedence over gitignore patterns NOT works for directories.
	// For performance reasons, we skip directories that are ignored by gitignore.
	// To avoid scanning the files in the ignored directories.
	// To analyze thus directory, you can specific the directory as one of root directories.
	//
	return isVCSDir(leaf) || (GitIgnore != nil && GitIgnore.MatchesPath(cutRootDirPath))
}

func ShouldIgnoreFile(cutRootDirPath string) bool {
	return GitIgnore != nil && GitIgnore.MatchesPath(cutRootDirPath)
}

// walkDir recursively descends path, calling walkDirFn.
func walkDir(path string, d fs.DirEntry, walkDirFn fs.WalkDirFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	if shouldIgnoreDir(path, d) {
		return
	}

	// First call, to report the path itself: file or dir
	if err := walkDirFn(path, d, nil); err != nil {
		log.Println("walkDirFn error:", err)
		return
	}

	// Only read directories.
	if !d.IsDir() {
		return
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		// Second call, to report ReadDir error.
		err = walkDirFn(path, d, err)
		if err != nil {
			log.Println("walkDirFn error when readding dir:", err)
			return
		}
	}

	for _, d1 := range dirs {
		path1 := filepath.Join(path, d1.Name())
		d1 := d1
		wg.Add(1)
		go walkDir(path1, d1, walkDirFn, wg)
	}
}
