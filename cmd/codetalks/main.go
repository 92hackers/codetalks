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

	"github.com/92hackers/codetalks/internal"
	"github.com/92hackers/codetalks/internal/file"
	"github.com/92hackers/codetalks/internal/language"
	"github.com/92hackers/codetalks/internal/output"
	"github.com/92hackers/codetalks/internal/scanner"
	"github.com/92hackers/codetalks/internal/utils"
	"github.com/92hackers/codetalks/internal/view_mode"
)

type cliOptions struct {
	isDebug       bool
	isProfile     bool
	outputFormat  string
	viewMode      string
	match         string // regular expression to match
	isShowMatched bool
	ignore        string // regular expression to ignore
	isShowIgnored bool
}

func parseOptions() *cliOptions {
	// Cli flags processing
	isPrintVersion := flag.Bool("version", false, "Print the version of the codetalks")
	isDebug := flag.Bool("debug", false, "Enable debug mode (display internal analyze logs)")
	isProfile := flag.Bool("profile", false, "Enable profile mode")

	outputFormat := flag.String("output", output.OutputFormatTable, "Output format of the codetalks")
	viewMode := flag.String("view", view_mode.ViewModeOverview, "View mode of the codetalks")

	match := flag.String("match", "", "Only analyze files or directories that match the regular expression")
	isShowMatched := flag.Bool("show-matched", false, "Show matched files (works with -match option)")

	ignore := flag.String("ignore", "", "Ignore files or directories that match the regular expression")
	isShowIgnored := flag.Bool("show-ignored", false, "Show ignored files (works with -ignore option)")

	// Parse the flags
	flag.Parse()

	if *isPrintVersion {
		fmt.Println("CodeTalks " + Version)
		return nil
	}

	return &cliOptions{
		match:         *match,
		isShowMatched: *isShowMatched,
		ignore:        *ignore,
		isShowIgnored: *isShowIgnored,
		isDebug:       *isDebug,
		isProfile:     *isProfile,
		outputFormat:  *outputFormat,
		viewMode:      *viewMode,
	}
}

// Init global options with cli options
func inigGlobalOpts(cliOptions *cliOptions, rootDirs []string) *internal.GlobalOptions {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &internal.GlobalOptions{
		MatchRegexp:      cliOptions.match,
		IgnoreRegexp:     cliOptions.ignore,
		IsShowMatched:    cliOptions.isShowMatched,
		IsShowIgnored:    cliOptions.isShowIgnored,
		IsDebugEnabled:   cliOptions.isDebug,
		ViewMode:         cliOptions.viewMode,
		OutputFormat:     cliOptions.outputFormat,
		IsProfileEnabled: cliOptions.isProfile,
		RootDirs:         rootDirs,
		Cwd:              cwd,
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

	// Initialize global options
	internal.GlobalOpts = inigGlobalOpts(cliOptions, rootDirs)

	// Record the time consumed
	defer utils.AnalyzeTimeConsumed()()

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
	scanner.Config(cliOptions.match, cliOptions.ignore)
	scanner.Scan(rootDirs)

	// Analyze code files
	if internal.GlobalOpts.IsDebugEnabled {
		fmt.Println("rootDirs:", rootDirs)
		fmt.Println("isDebug:", cliOptions.isDebug)
		fmt.Println("output format:", cliOptions.outputFormat)
		fmt.Println("view mode:", cliOptions.viewMode)
		fmt.Println("AllCodeFiles:", len(file.AllCodeFiles))
		utils.TimeIt(language.AnalyzeAllLanguages)
	} else {
		language.AnalyzeAllLanguages()
	}

	// Set the view mode
	// view_mode.SetViewMode(cliOptions.viewMode)

	// Output the result
	output.Output(cliOptions.outputFormat)
}
