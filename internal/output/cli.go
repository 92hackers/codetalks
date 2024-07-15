/**
Output on the console as table.
*/

package output

import (
	"fmt"
	"golang.org/x/text/message"
	"sort"

	"github.com/92hackers/codetalks/internal/language"
)

var rowLength = 80
var tableLine = "==============================================================================="
var printer = message.NewPrinter(message.MatchLanguage("en"))

func OutputCliTable() {
	renderHeader()
	renderData()
	renderFooter()
}

func renderHeader() {
	tableHeader := "| Language     | Files     | Total      | Comments   | Blanks    | Code       |"
	fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)
	fmt.Printf("%.[2]*[1]s\n", tableHeader, rowLength)
	fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)
}

func renderData() {
	// Default sort criteria.
	sortLanguagesByCode()

	for _, lang := range language.AllLanguages {
		printer.Printf("| %-12s | %-9d | %-10d | %-10d | %-9d | %-10d |\n", lang.Name, lang.FileCount, lang.TotalLines, lang.CommentCount, lang.BlankCount, lang.CodeCount)
	}
}

func renderFooter() {
	if len(language.AllLanguages) > 0 {
		fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)
	}

	aggregateStats := language.AllLanguageAggregateStats
	printer.Printf("| %-12s | %-9d | %-10d | %-10d | %-9d | %-10d |\n", "Total", aggregateStats.TotalFiles, aggregateStats.TotalLines, aggregateStats.TotalComment, aggregateStats.TotalBlank, aggregateStats.TotalCode)

	fmt.Printf("%.[2]*[1]s\n", tableLine, rowLength)
}

// Sort by total lines of code in descending order
func sortLanguagesByCode() {
	sort.Slice(language.AllLanguages, func(i, j int) bool {
		a := language.AllLanguages[i]
		b := language.AllLanguages[j]

		// Sort by code count firstly
		if a.CodeCount != b.CodeCount {
			return a.CodeCount > b.CodeCount
		}

		// Sort by total lines if code count is same
		if a.TotalLines != b.TotalLines {
			return a.TotalLines > b.TotalLines
		}

		// Sort by files count if total lines are same
		if a.FileCount != b.FileCount {
			return a.FileCount > b.FileCount
		}

		return true
	})
}
