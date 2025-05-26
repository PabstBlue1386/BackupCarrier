package fileutils

import (
	"os"
)

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == os.ErrNotExist {
		return false, err
	} else {
		return true, err
	}
}
