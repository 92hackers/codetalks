/**
Output on the console as table.
*/

package output

import (
	"fmt"
	// "github.com/92hackers/code-talks/internal/language"
)

func OutputCliTable() {
	rowLength := 80
	tableLine := "==============================================================================="
	tableHeader := "| Language        | Files     | Total      | Comments       | Blanks   | Code      |"

	fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)
	fmt.Printf("%.[2]*[1]s\n", tableHeader, rowLength)
	fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)
}
