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
	"time"

	"github.com/92hackers/code-talks/internal"
	"github.com/92hackers/code-talks/internal/file"
	"github.com/92hackers/code-talks/internal/language"
	"github.com/92hackers/code-talks/internal/output"
	"github.com/92hackers/code-talks/internal/scanner"
	"github.com/92hackers/code-talks/internal/utils"
	"github.com/92hackers/code-talks/internal/view_mode"
)

type cliOptions struct {
	isDebug      bool
	isProfile    bool
	outputFormat string
	viewMode     string
}

func parseOptions() *cliOptions {
	// Cli flags processing
	isPrintVersion := flag.Bool("version", false, "Print the version of the code-talks")
	isDebug := flag.Bool("debug", false, "Enable debug mode")
	isProfile := flag.Bool("profile", false, "Enable profile mode")

	outputFormat := flag.String("output", output.OutputFormatTable, "Output format of the code-talks")
	viewMode := flag.String("view", view_mode.ViewModeOverview, "View mode of the code-talks")

	// Parse the flags
	flag.Parse()

	if *isPrintVersion {
		fmt.Println("CodeTalks v0.1")
		return nil
	}

	return &cliOptions{
		isDebug:      *isDebug,
		isProfile:    *isProfile,
		outputFormat: *outputFormat,
		viewMode:     *viewMode,
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

	// Record the time consumed
	start := time.Now()
	defer func() {
		fmt.Println("Analyze time consumed: ", time.Since(start))
	}()

	// Profile the code
	if cliOptions.isProfile {
		if cpuDone := utils.StartCPUProfile(); cpuDone != nil {
			defer cpuDone()
		}
		if memDone := utils.StartMemoryProfile(); memDone != nil {
			defer memDone()
		}
		if traceDone := utils.StartTrace(); traceDone != nil {
			defer traceDone()
		}
	}

	// Scan root directory
	scanner.Scan(rootDir)

	if internal.IsDebugEnabled {
		fmt.Println("isDebug: ", cliOptions.isDebug)
		fmt.Println("output format: ", cliOptions.outputFormat)
		fmt.Println("view mode: ", cliOptions.viewMode)
		fmt.Println("rootDir: ", rootDir)
		fmt.Println("AllCodeFiles: ", len(file.AllCodeFiles))

		// Analyze code files
		utils.TimeIt(language.AnalyzeAllLanguages)
	} else {
		language.AnalyzeAllLanguages()
	}

	// Set the view mode
	// view_mode.SetViewMode(cliOptions.viewMode)

	// Output the result
	output.Output(cliOptions.outputFormat)
}
