/**

Scanner.

*/

package scanner

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func handler(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	// dir
	if d.IsDir() {
		fmt.Println("dir:", path)
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
