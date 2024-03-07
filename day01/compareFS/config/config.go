package config

import (
	"errors"
	"flag"
	"fmt"
)

func GetSnapshots() (oldSnapshot string, newSnapshot string, err error) {
	var NoFilenameError = errors.New("no such file")

	flag.StringVar(&oldSnapshot, "old", "", "Old snapshot file name")
	flag.StringVar(&newSnapshot, "new", "", "New shapshot file name")
	flag.Parse()

	if oldSnapshot == "" || newSnapshot == "" {
		fmt.Println(NoFilenameError)
		return "", "", NoFilenameError
	}
	return oldSnapshot, newSnapshot, nil
}
