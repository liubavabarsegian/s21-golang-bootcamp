package config

import (
	"errors"
	"flag"
)

var NoFilenameError = errors.New("No such file")

func CheckDBFileName() (OldDB string, NewDB string, err error) {
	flag.StringVar(&OldDB, "old", "", "")
	flag.StringVar(&NewDB, "new", "", "")
	flag.Parse()

	if OldDB == "" || NewDB == "" {
		return "", "", NoFilenameError
	}
	return
}
