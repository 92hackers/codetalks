/**
Test walkdir functions
*/

package utils

import (
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsVCSDir(t *testing.T) {
	vscDirs := []string{".git", ".svn", ".hg", ".bzr", ".cvs"}
	for _, dir := range vscDirs {
		AssertEqual(t, isVCSDir(dir), true)
	}
	assert.True(t, isVCSDir(".git"))
	assert.True(t, isVCSDir(".svn"))
	assert.True(t, isVCSDir(".hg"))
	assert.True(t, isVCSDir(".bzr"))
	assert.True(t, isVCSDir(".cvs"))
	assert.False(t, isVCSDir(".fakegit"))
}

func clearState() {
	GitIgnore = nil
	CurrentRootDir = ""
}

func TestWalkDir(t *testing.T) {
	assert.Nil(t, GitIgnore)
	assert.Empty(t, CurrentRootDir)

	root := filepath.Join("..", "..", "testdata/gitignore")
	count := 0
	// Test WalkDir
	err := WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		t.Logf("path: %s", path)
		count++
		return nil
	})

	assert.Nil(t, err)
	assert.NotNil(t, GitIgnore)
	assert.NotEmpty(t, CurrentRootDir)
	assert.Equal(t, count, 7)

	t.Cleanup(clearState)
}
