package vedautils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	VedafileName  = "Veda.yaml"
	OverwriteFile = ".guru.yaml"
)

const sep = string(filepath.Separator)

const defaultVedafile = `# Default values for %s.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
name: blubb
`

func Create(name, dir string) (string, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
		return path, err
	}

	vDir := filepath.Join(path, name)

	files := []struct {
		path    string
		content []byte
	}{
		{
			path:    filepath.Join(vDir, VedafileName),
			content: []byte(fmt.Sprintf(defaultVedafile, VedafileName)),
		},
	}

	for _, file := range files {
		// TODO: Check wether file already exists

		if err := os.MkdirAll(filepath.Dir(file.path), 0755); err != nil {
			return file.path, err
		}
		if err := ioutil.WriteFile(file.path, file.content, 0644); err != nil {
			return file.path, err
		}
	}

	return vDir, err
}
