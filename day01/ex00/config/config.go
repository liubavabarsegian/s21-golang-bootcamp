package config

import (
	"errors"
	"flag"
)

var NoFilenameError = errors.New("Передайте название файла через флаг -f")

func CheckDBFileName() (DBFileName string, err error) {
	flag.StringVar(&DBFileName, "f", "", "DataBase file name")
	flag.Parse()

	if DBFileName == "" {
		return "", NoFilenameError
	}
	return DBFileName, err
}
