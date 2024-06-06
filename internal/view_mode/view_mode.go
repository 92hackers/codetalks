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
	ViewModeDirectoriesDepth = "directories_depth"
)

// Default view mode: overview, nospecific processing
func SetViewMode(viewMode string) {
	fmt.Println("View mode:", viewMode)
}
