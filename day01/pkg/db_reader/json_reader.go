package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type JSONReader struct {
	cakes Recipes
}

type JSONconverter struct {
}

func (reader JSONReader) ReadDB(filename string) (Recipes, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return reader.cakes, err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &reader.cakes)
	return reader.cakes, err
}

func (converter JSONconverter) Convert(cakes Recipes) {
	xmlCakes, _ := xml.MarshalIndent(cakes, " ", "    ")
	fmt.Println(string(xmlCakes))
}
