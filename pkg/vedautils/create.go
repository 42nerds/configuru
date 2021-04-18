package vedautils

import "path/filepath"

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
