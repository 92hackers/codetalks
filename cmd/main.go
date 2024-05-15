/**

CodeTalks command

*/

package main

import (
	"flag"
	"fmt"
	"github.com/92hackers/code-talks/internal/scanner"
	"log"
	"os"
	"path/filepath"
)

type cliOptions struct {
	outputFormat string
}

func parseOptions() *cliOptions {
	// Cli flags processing
	isPrintVersion := flag.Bool("version", false, "Print the version of the code-talks")
	outputFormat := flag.String("output", "table", "Output format of the code-talks")
	flag.Parse()

	if *isPrintVersion {
		fmt.Println("CodeTalks v0.1")
		return nil
	}

	return &cliOptions{
		outputFormat: *outputFormat,
	}
}

func getRootDir() string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	rootDir := workingDir

	cliArgs := flag.Args()
	if len(cliArgs) > 0 {
		dir := cliArgs[0] // Only support analyzing one directory currently.
		rootDir = filepath.Join(workingDir, dir)

		dirInfo, err := os.Stat(rootDir) // Follow the symbolic link
		switch {
		case os.IsNotExist(err):
			log.Fatalf("Directory %s does not exist", rootDir)
		case os.IsPermission(err):
			log.Fatalf("No permission to access directory %s", rootDir)
		case err != nil:
			log.Fatalf("Error accessing directory %s", rootDir)
		}

		if !dirInfo.IsDir() {
			log.Fatalf("%s is not a directory, only directory supported now.", rootDir)
		}
	}

	return rootDir
}

func main() {
	// Parse the cli options
	cliOptions := parseOptions()
	fmt.Println("output format: ", cliOptions.outputFormat)

	// Get the root directory
	rootDir := getRootDir()
	fmt.Println("rootDir: ", rootDir)

	// Scan the directory
	scanner.Scan(rootDir)
}
