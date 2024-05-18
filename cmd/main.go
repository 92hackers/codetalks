/**

CodeTalks command

*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/92hackers/code-talks/internal"
	"github.com/92hackers/code-talks/internal/file"
	"github.com/92hackers/code-talks/internal/language"
	"github.com/92hackers/code-talks/internal/output"
	"github.com/92hackers/code-talks/internal/scanner"
	"github.com/92hackers/code-talks/internal/utils"
)

type cliOptions struct {
	isDebug      bool
	outputFormat string
}

func parseOptions() *cliOptions {
	// Cli flags processing
	isPrintVersion := flag.Bool("version", false, "Print the version of the code-talks")
	outputFormat := flag.String("output", output.OutputFormatTable, "Output format of the code-talks")
	isDebug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if *isPrintVersion {
		fmt.Println("CodeTalks v0.1")
		return nil
	}

	return &cliOptions{
		isDebug:      *isDebug,
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
		if rootDir, err = filepath.Abs(dir); err != nil {
			log.Fatalf("Error getting absolute path of %s", dir)
		}

		// Check if the directory exists
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

	// Set the debug flag
	internal.IsDebugEnabled = cliOptions.isDebug

	// Get the root directory
	rootDir := getRootDir()

	// Scan root directory
	scanner.Scan(rootDir)

	if internal.IsDebugEnabled {
		fmt.Println("isDebug: ", cliOptions.isDebug)
		fmt.Println("output format: ", cliOptions.outputFormat)
		fmt.Println("rootDir: ", rootDir)
		fmt.Println("AllCodeFiles: ", len(file.AllCodeFiles))

		// Analyze code files
		utils.TimeIt(language.AnalyzeAllLanguages)
	} else {
		language.AnalyzeAllLanguages()
	}
	// Slow version
	// utils.TimeIt(language.AnalyzeAllLanguagesSlow)

	// Output the result
	output.Output(cliOptions.outputFormat)

	// for k, v := range language.AllLanguagesMap {
	// 	fmt.Println(k, ": ")
	//
	// 	// Stats for every language.
	// 	fmt.Println("FileCount: ", v.FileCount)
	// 	fmt.Println("TotalLines: ", v.TotalLines)
	// 	fmt.Println("CodeCount: ", v.CodeCount)
	// 	fmt.Println("CommentCount: ", v.CommentCount)
	// 	fmt.Println("BlankCount: ", v.BlankCount)
	// 	fmt.Println("-=-=-=-=-=-=-")
	//
	// 	// Stats for every file.
	// 	// for _, codeFile := range v.CodeFiles {
	// 	// 	fmt.Print(codeFile.Path, " ")
	// 	// 	fmt.Println("File size: ", codeFile.Size)
	// 	// 	fmt.Println("TotalLines: ", codeFile.TotalLines)
	// 	// 	fmt.Println("CodeCount: ", codeFile.CodeCount)
	// 	// 	fmt.Println("CommentCount: ", codeFile.CommentCount)
	// 	// 	fmt.Println("BlankCount: ", codeFile.BlankCount)
	// 	// 	fmt.Println("-=-=-=-=-=-=-")
	// 	// }
	// }
}
