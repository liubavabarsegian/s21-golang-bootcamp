package dbreader

import (
	"day01/dbReader/config"
	"encoding/xml"
	"fmt"
	"path/filepath"
)

type DBReader interface {
	ReadDB(filename string) (Cakes, error)
}

type DBConverter interface {
	Convert(cakes Cakes)
}

type Cakes struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

func Call() {
	filename, err := config.CheckDBFileName()
	if err != nil {
		fmt.Println(err)
		return
	}
	ConvertFile(filename)
}

func GetCakes(filename string) (cakes Cakes, err error) {
	fileExtension := filepath.Ext(filename)
	var reader DBReader
	switch fileExtension {
	case ".json":
		reader = JSONReader{}
	case ".xml":
		reader = XMLReader{}
	default:
		fmt.Println("Unsupported file format:", fileExtension)
		return
	}
	cakes, err = reader.ReadDB(filename)
	return cakes, err
}

func ConvertFile(filename string) {
	fileExtension := filepath.Ext(filename)
	var converter DBConverter
	switch fileExtension {
	case ".json":
		converter = JSONconverter{}
	case ".xml":
		converter = XMLconverter{}
	default:
		fmt.Println("Unsupported file format:", fileExtension)
		return
	}
	cakes, err := GetCakes(filename)
	if err == nil {
		converter.Convert(cakes)
	}
}
