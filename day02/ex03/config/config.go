package config

import (
	"errors"
	"flag"
	"os"
)

var ErrNoSuchDir = errors.New("no such directory")
var ErrNoFilesPassed = errors.New("no files passed")

func GetArgs() (directory string, filenames []string, err error) {
	flag.StringVar(&directory, "a", "", "Directory to archive to")
	flag.Parse()

	if directory != "" {
		fileInfo, err := os.Stat(directory)
		if err != nil {
			return "", nil, err
		}

		if !fileInfo.IsDir() {
			err = ErrNoSuchDir
			return "", nil, err
		}
	}

	filenames = flag.Args()
	if len(filenames) == 0 {
		err = ErrNoFilesPassed
		return
	}

	return
}
