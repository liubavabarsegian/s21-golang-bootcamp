package dbreader

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"readDB/config"
)

type DBReader interface {
	readDB(filename string) (Cakes, error)
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

func Read() {
	filename, err := config.CheckDBFileName()
	if err != nil {
		fmt.Println(err)
	}
	fileExtension := filepath.Ext(filename)
	var reader DBReader
	var converter DBConverter
	switch fileExtension {
	case ".json":
		reader = JSONReader{}
		converter = JSONconverter{}
	case ".xml":
		reader = XMLReader{}
		converter = XMLconverter{}
	}
	cakes, _ := reader.readDB(filename)
	converter.Convert(cakes)
}
