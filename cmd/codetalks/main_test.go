/**

Testing for codetalks binary

*/

package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	// "github.com/92hackers/codetalks/internal/utils"
)

// Because we have already cd into the root directory,
// we can use the relative path to the root directory.
var (
	codeTalksBinary = filepath.Join("bin", "codetalks")
	normalCodebase  = filepath.Join("testdata", "normal")
)
var cwd string

func execCommand(command string, args ...string) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), err
}

func init() {
	os.Chdir(filepath.Join("..", ".."))
	_, err := execCommand("make", "build") // Build codetaalks binary
	if err != nil {
		panic(err)
	}
	cwd, _ = os.Getwd()
}

func TestCMD(t *testing.T) {
	cmd := exec.Command(filepath.Join(cwd, codeTalksBinary), filepath.Join(cwd, normalCodebase))
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("❌ Error: %s\nOutput: %s\n", err, output)
	}
	expected := `===============================================================================
| Language     | Files     | Total      | Comments   | Blanks    | Code       |
===============================================================================
| Rust         | 1         | 48         | 10         | 5         | 33         |
| C            | 1         | 50         | 9          | 9         | 32         |
| HTML         | 1         | 46         | 6          | 8         | 32         |
| Go           | 2         | 51         | 10         | 10        | 31         |
| YAML         | 1         | 34         | 3          | 2         | 29         |
| Markdown     | 1         | 41         | 0          | 15        | 26         |
| Shell        | 1         | 45         | 14         | 8         | 23         |
| C#           | 1         | 26         | 9          | 3         | 14         |
| Python       | 1         | 15         | 5          | 3         | 7          |
| Plain Text   | 1         | 1          | 0          | 0         | 1          |
===============================================================================
| Total        | 11        | 357        | 66         | 63        | 228        |
===============================================================================
`
	if !strings.HasPrefix(string(output), expected) {
		t.Errorf("❌ Expected %s but got %s", expected, string(output))
	}
}
