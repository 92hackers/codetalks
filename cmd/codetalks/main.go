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

	"github.com/92hackers/codetalks/internal"
	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
	"github.com/92hackers/codetalks/internal/output"
	"github.com/92hackers/codetalks/internal/scanner"
	"github.com/92hackers/codetalks/internal/utils"
	"github.com/92hackers/codetalks/internal/view_mode"
)

type cliOptions struct {
	isDebug      bool
	isProfile    bool
	outputFormat string
	viewMode     string
	re           string // regular expression
	ignore       string // ignore regular expression
}

func parseOptions() *cliOptions {
	// Cli flags processing
	isPrintVersion := flag.Bool("version", false, "Print the version of the codetalks")
	isDebug := flag.Bool("debug", false, "Enable debug mode")
	isProfile := flag.Bool("profile", false, "Enable profile mode")

	outputFormat := flag.String("output", output.OutputFormatTable, "Output format of the codetalks")
	viewMode := flag.String("view", view_mode.ViewModeOverview, "View mode of the codetalks")
	re := flag.String("re", "", "Only analyze files or directories that match the regular expression")
	ignore := flag.String("ignore", "", "Ignore files or directories that match the regular expression")

	// Parse the flags
	flag.Parse()

	if *isPrintVersion {
		fmt.Println("CodeTalks " + Version)
		return nil
	}

	return &cliOptions{
		re:           *re,
		ignore:       *ignore,
		isDebug:      *isDebug,
		isProfile:    *isProfile,
		outputFormat: *outputFormat,
		viewMode:     *viewMode,
	}
}

func getRootDirs() []string {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cliArgs := flag.Args()

	if len(cliArgs) == 0 {
		return []string{workingDir}
	}

	var rootDirs []string
	uniqueDirSet := utils.NewSet()

	// Multiple directories can be provided
	for _, dir := range cliArgs {
		rootDir, err := filepath.Abs(dir)
		if err != nil {
			log.Fatalf("Error getting absolute path of: %s", dir)
		}

		// TODO: remove subdirectory of another directory.
		// Duplicate check
		if uniqueDirSet.Contains(rootDir) {
			continue
		}

		// Check if the directory or file exists
		_, err = os.Stat(rootDir) // Follow the symbolic link, but WalkDir() does not follow symbolic links.
		switch {
		case os.IsNotExist(err):
			log.Fatalf("Directory %s does not exist", rootDir)
		case os.IsPermission(err):
			log.Fatalf("No permission to access directory %s", rootDir)
		case err != nil:
			log.Fatalf("Error accessing directory %s", rootDir)
		}

		// Files also can be appended to the rootDirs
		rootDirs = append(rootDirs, rootDir)
		uniqueDirSet.Add(rootDir)
	}

	return rootDirs
}

func formatDuration(d time.Duration) string {
	scale := 100 * time.Second
	// look for the max scale that is smaller than d
	for scale > d {
		scale = scale / 10
	}
	return d.Round(scale / 100).String()
}

func main() {
	// Parse the cli options
	cliOptions := parseOptions()
	if cliOptions == nil {
		os.Exit(0)
	}

	// Set the debug flag
	internal.IsDebugEnabled = cliOptions.isDebug

	// Get the root directories
	rootDirs := getRootDirs()

	// Record the time consumed
	start := time.Now()
	defer func() {
		fmt.Println("Analyze time consumed: ", formatDuration(time.Since(start)))
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

	// Scan root directories
	scanner.Config(cliOptions.re, cliOptions.ignore)
	scanner.Scan(rootDirs)

	if internal.IsDebugEnabled {
		fmt.Println("rootDirs: ", rootDirs)
		fmt.Println("isDebug: ", cliOptions.isDebug)
		fmt.Println("output format: ", cliOptions.outputFormat)
		fmt.Println("view mode: ", cliOptions.viewMode)
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
