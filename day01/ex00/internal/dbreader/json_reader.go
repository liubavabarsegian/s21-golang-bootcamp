package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type JSONReader struct {
	cakes Cakes
}

type JSONconverter struct {
}

func (reader JSONReader) readDB(filename string) (Cakes, error) {
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()
	}

	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &reader.cakes)
	return reader.cakes, err
}

func (converter JSONconverter) Convert(cakes Cakes) {
	xml_cakes, _ := xml.MarshalIndent(cakes, " ", "    ")
	fmt.Println(string(xml_cakes))
}
