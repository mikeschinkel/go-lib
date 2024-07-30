package lib

import (
	"errors"
	"io/fs"
	"os"
)

func MustIsDir(path string) (isDir bool) {
	isDir, err := IsDir(path)
	if err != nil {
		panic(err)
	}
	return isDir
}

func IsDir(path string) (isDir bool, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(path)
	switch {
	case err == nil:
		isDir = fi.IsDir()
	case errors.Is(err, fs.ErrNotExist):
		break
	default:
		panic(err)
	}
	return isDir, err
}
