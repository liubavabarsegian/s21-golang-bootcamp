package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type XMLReader struct {
	cakes Recipes
}

type XMLtoJSON struct {
}

type XMLconverter struct {
}

func (reader XMLReader) ReadDB(filename string) (Recipes, error) {
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()
	}

	byteValue, _ := io.ReadAll(file)
	xml.Unmarshal(byteValue, &reader.cakes)
	return reader.cakes, err
}

func (converter XMLconverter) Convert(cakes Recipes) {
	jsonCakes, _ := json.MarshalIndent(cakes, "", "    ")
	fmt.Println(string(jsonCakes))
}
