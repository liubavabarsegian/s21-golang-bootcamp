package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type XMLReader struct {
	cakes Cakes
}

type XMLtoJSON struct {
}

type XMLconverter struct {
}

func (reader XMLReader) ReadDB(filename string) (Cakes, error) {
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()
	}

	byteValue, _ := io.ReadAll(file)
	xml.Unmarshal(byteValue, &reader.cakes)
	return reader.cakes, err
}

func (converter XMLconverter) Convert(cakes Cakes) {
	json_cakes, _ := json.MarshalIndent(cakes, "", "    ")
	fmt.Println(string(json_cakes))
}
