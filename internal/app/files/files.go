package files

import (
	"os"
	"strings"
)

func GetDir(path string) (string, error) {
	path, err := FormatDir(path)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(path, "/") {
		return path, nil
	}
	workdir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return workdir + "/" + path, nil
}

func FormatDir(dir string) (string, error) {
	if strings.HasSuffix(dir, "/") {
		return dir, nil
	} else {
		return dir + "/", nil
	}
}
