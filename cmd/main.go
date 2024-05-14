/**

CodeTalks command

*/

package main

import (
	"fmt"
	"github.com/92hackers/code-talks/internal/file"
	// "os"
)

func main() {
	file := &file.CodeTalksFile{
		Code:     32,
		Comments: 10,
		Blanks:   5,
		Name:     "main.go",
		Path:     "main.go",
		Language: "Go",
	}

	fmt.Println(file.Language)

	fmt.Println(file.GetCode())
}
