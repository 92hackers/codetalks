/**
View mode
*/

package view_mode

import (
	"fmt"
)

const (
	ViewModeOverview         = "overview"
	ViewModeFiles            = "files"
	ViewModeDirectories      = "directories"
	ViewModeDirectoriesDepth = "directories_depth" // Directories with depth
)

func ValidateViewMode(viewMode string) string {
	switch viewMode {
	case ViewModeOverview, ViewModeFiles, ViewModeDirectories:
		return viewMode
	default:
		panic("Valid view modes are: overview, files, directories")
	}
}

// Default view mode: overview, nospecific processing
func SetViewMode(viewMode string) {
	fmt.Println("View mode:", viewMode)
}
