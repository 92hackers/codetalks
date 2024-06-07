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
	_, err := execCommand("make", "build") // Build codetalks binary
	if err != nil {
		panic(err)
	}
	cwd, _ = os.Getwd()
}

func expectOutputWithCommand(t *testing.T, expected string, args ...string) string {
	t.Helper()
	cmd := exec.Command(filepath.Join(cwd, codeTalksBinary), args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("❌ Error: %s\nOutput: %s\n", err, output)
	}
	outputStr := string(output)
	if !strings.Contains(outputStr, expected) {
		t.Log("Actual command:", cmd)
		t.Errorf("❌ Expected %s but got %s", expected, outputStr)
	}
	return outputStr
}

func TestCMD(t *testing.T) {
	expected := `===============================================================================
| Language     | Files     | Total      | Comments   | Blanks    | Code       |
===============================================================================
| Rust         | 1         | 48         | 10         | 5         | 33         |
| C            | 1         | 50         | 9          | 9         | 32         |
| HTML         | 1         | 46         | 6          | 8         | 32         |
| Go           | 2         | 51         | 10         | 10        | 31         |
| YAML         | 1         | 34         | 3          | 2         | 29         |
| Markdown     | 1         | 41         | 0          | 15        | 26         |
| Vue          | 1         | 36         | 9          | 3         | 24         |
| Shell        | 1         | 45         | 14         | 8         | 23         |
| C#           | 1         | 26         | 9          | 3         | 14         |
| Python       | 1         | 15         | 5          | 3         | 7          |
| Plain Text   | 1         | 1          | 0          | 0         | 1          |
===============================================================================
| Total        | 12        | 393        | 75         | 66        | 252        |
===============================================================================
`
	expectOutputWithCommand(t, expected, filepath.Join(cwd, normalCodebase))
}

func TestCMDWithMatchOption(t *testing.T) {
	expected := `===============================================================================
| Language     | Files     | Total      | Comments   | Blanks    | Code       |
===============================================================================
| Rust         | 1         | 48         | 10         | 5         | 33         |
| Go           | 2         | 51         | 10         | 10        | 31         |
| Python       | 1         | 15         | 5          | 3         | 7          |
===============================================================================
| Total        | 4         | 114        | 25         | 18        | 71         |
===============================================================================
`
	codebase := filepath.Join(cwd, normalCodebase)
	expectOutputWithCommand(t, expected, "-match", ".rs$ .go$ .py$ ", codebase)

	t.Run("TestCMDWithMatchedShowed", func(t *testing.T) {
		output := expectOutputWithCommand(t, expected, "-match", ".rs$ .go$ .py$ ", "--show-matched", codebase)
		matchStr := "File matched"
		if !strings.HasPrefix(output, matchStr) {
			t.Errorf("❌ %s not found in %s", matchStr, output)
		}
	})
}

func TestCMDWithIgnoreOption(t *testing.T) {
	expected := `===============================================================================
| Language     | Files     | Total      | Comments   | Blanks    | Code       |
===============================================================================
| C            | 1         | 50         | 9          | 9         | 32         |
| HTML         | 1         | 46         | 6          | 8         | 32         |
| YAML         | 1         | 34         | 3          | 2         | 29         |
| Markdown     | 1         | 41         | 0          | 15        | 26         |
| Vue          | 1         | 36         | 9          | 3         | 24         |
| Shell        | 1         | 45         | 14         | 8         | 23         |
| C#           | 1         | 26         | 9          | 3         | 14         |
| Plain Text   | 1         | 1          | 0          | 0         | 1          |
===============================================================================
| Total        | 8         | 279        | 50         | 48        | 181        |
===============================================================================
`
	codebase := filepath.Join(cwd, normalCodebase)
	expectOutputWithCommand(t, expected, "-ignore", ".rs$ .go$ .py$ ", codebase)

	t.Run("TestCMDWithIgnoredShowed", func(t *testing.T) {
		output := expectOutputWithCommand(t, expected, "-ignore", ".rs$ .go$ .py$ ", "--show-ignored", codebase)
		matchStr := "File ignored"
		if !strings.HasPrefix(output, matchStr) {
			t.Errorf("❌ %s not found in %s", matchStr, output)
		}
	})
}
