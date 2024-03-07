package config

import (
	"errors"
	"flag"
)

func GetSnapshots() (oldSnapshot string, newSnapshot string, err error) {
	var NoFilenameError = errors.New("передайте файлы в формате: --old snapshot1 --new snapshot")

	flag.StringVar(&oldSnapshot, "old", "", "Old snapshot file name")
	flag.StringVar(&newSnapshot, "new", "", "New shapshot file name")
	flag.Parse()

	if oldSnapshot == "" || newSnapshot == "" {
		return "", "", NoFilenameError
	}
	return oldSnapshot, newSnapshot, nil
}
