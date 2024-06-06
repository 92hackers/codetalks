/**

Testing for view mode package

*/

package view_mode

import (
	"testing"

	"github.com/92hackers/codetalks/internal/utils"
)

func TestSetViewMode(t *testing.T) {
	output := utils.CaptureStdout(func() {
		SetViewMode(ViewModeOverview)
	})
	utils.AssertEqual(t, output, "View mode: overview\n")
}
