/**
Output on the console as table.
*/

package output

import (
	"fmt"
	"github.com/92hackers/code-talks/internal/language"
	"sort"
)

func OutputCliTable() {
	rowLength := 80
	tableLine := "==============================================================================="
	tableHeader := "| Language        | Files     | Total      | Comments       | Blanks   | Code      |"

	fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)
	fmt.Printf("%.[2]*[1]s\n", tableHeader, rowLength)
	fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)

	// Default sort criteria.
	sortLanguagesByCode()

	for _, lang := range language.AllLanguages {
		fmt.Printf("| %-15s | %-9d | %-10d | %-14d | %-8d | %-9d |\n", lang.Name, lang.FileCount, lang.TotalLines, lang.CommentCount, lang.BlankCount, lang.CodeCount)
	}
}

// Sort by total lines of code in descending order
func sortLanguagesByCode() {
	sort.Slice(language.AllLanguages, func(i, j int) bool {
		return language.AllLanguages[i].CodeCount > language.AllLanguages[j].CodeCount
	})
}
