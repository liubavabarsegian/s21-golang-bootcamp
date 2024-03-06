package config

import (
	"errors"
	"flag"
	"fmt"
)

func GetFiles() (OldDB string, NewDB string, err error) {
	var NoFilenameError = errors.New("no such file")

	flag.StringVar(&OldDB, "old", "", "Old database file name")
	flag.StringVar(&NewDB, "new", "", "New database file name")
	flag.Parse()

	if OldDB == "" || NewDB == "" {
		fmt.Println(NoFilenameError)
		return "", "", NoFilenameError
	}
	return OldDB, NewDB, nil
}
