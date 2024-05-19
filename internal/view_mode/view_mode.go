/**
View mode
*/

package view_mode

import (
	"fmt"
)

const ViewModeOverview = "overview"
const ViewModeFiles = "files"
const ViewModeDirectories = "directories"
const ViewModeDirectoriesDepth = "directories_depth"

func SetViewMode(viewMode string) {
	fmt.Println("View mode:", viewMode)
}
