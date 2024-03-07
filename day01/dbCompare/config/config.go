package config

import (
	"errors"
	"flag"
)

func GetFiles() (OldDB string, NewDB string, err error) {
	var NoFilenameError = errors.New("передайте файлы в формате: --old filename --new filename")

	flag.StringVar(&OldDB, "old", "", "Old database file name")
	flag.StringVar(&NewDB, "new", "", "New database file name")
	flag.Parse()

	if OldDB == "" || NewDB == "" {
		return "", "", NoFilenameError
	}
	return OldDB, NewDB, nil
}
