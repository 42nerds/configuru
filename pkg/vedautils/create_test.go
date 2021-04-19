package vedautils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "guru-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	_, err = Create("blubb", tmpDir)

	dir := filepath.Join(tmpDir, "blubb")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range []string{
		VedafileName,
	} {
		if _, err := os.Stat(filepath.Join(dir, file)); err != nil {
			t.Errorf("Expected %s file: %s", file, err)
		}
	}
	// Test if MkDirAll throws an error by trying to create a directory in root directory
	_, err = Create("blubb", "/non-existing-directory")
	assert.Error(t, err)
}
