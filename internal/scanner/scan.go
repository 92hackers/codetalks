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

	// dir
	if d.IsDir() {
		fmt.Println("dir:", path)

		// Skip VCS directories
		_, dir := filepath.Split(path)
		fmt.Println("last segment:", dir)
		if isVCSDir(dir) {
			return fs.SkipDir
		}

		return nil
	}

	// file
	fmt.Println("file:", path)

	return nil
}

func Scan(dir string) {
	// Scan directory
	err := filepath.WalkDir(dir, handler)
	if err != nil {
		log.Fatal(err)
	}
}
