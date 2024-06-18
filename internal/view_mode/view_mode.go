/**
View mode
*/

package view_mode

import (
	"fmt"
)

const (
	ViewModeOverview  = "overview"
	ViewModeFiles     = "files"
	ViewModeDirs      = "dirs"
	ViewModeDirsDepth = "dirs_depth" // Directories with depth
)

var (
	SubDirs []string
)

func ValidateViewMode(viewMode string) string {
	switch viewMode {
	case ViewModeOverview, ViewModeFiles, ViewModeDirs:
		return viewMode
	default:
		panic("Valid view modes are: overview, files, dirs")
	}
}

// Default view mode: overview, nospecific processing
func SetViewMode(viewMode string) {
	fmt.Println("View mode:", viewMode)
}
