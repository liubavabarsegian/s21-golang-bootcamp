package dbreader

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"readDB/config"
)

type DBReader interface {
	readDB(file *os.File) (Cakes, error)
}
type Cakes struct {
	Recipes xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count int    `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

func Read() {
	filename, err := config.CheckDBFileName()
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fileExtension := filepath.Ext(filename)
		var reader DBReader
		switch fileExtension {
		case ".json":
			reader = JSONReader{}
		case ".xml":
			reader = XMLReader{}
		}
		reader.readDB(file)
		defer file.Close()
	}
}
