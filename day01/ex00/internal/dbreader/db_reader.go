package dbreader

import (
	"fmt"
	"os"
	"path/filepath"
	"readDB/config"
	"readDB/internal/converter"
)

func Read() {
	filename, err := config.GetDBFileName()
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fileExtension := filepath.Ext(filename)
		switch fileExtension {
		case ".json":
			fmt.Println("successefully opened json file")
			converter.JSONtoXML(file)
		case ".xml":
			fmt.Println("successefully opened xml file")
		default:
			fmt.Println("wrong file format")
		}
		defer file.Close()
	}
}
