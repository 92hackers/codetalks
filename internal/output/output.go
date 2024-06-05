/**

Output package is used to output the analyze result in different formats.

*/

package output

const (
	OutputFormatTable string = "table"
	OutputFormatJSON  string = "json"
)

func Output(outputFormat string) {
	switch outputFormat {
	case OutputFormatTable:
		OutputCliTable()
	case OutputFormatJSON:
		OutputJSON()
	default:
		OutputCliTable()
	}
}
