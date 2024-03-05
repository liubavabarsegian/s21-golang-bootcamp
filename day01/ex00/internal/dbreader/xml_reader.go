package dbreader

import (
	"fmt"
	"os"
)

type XMLReader struct {
	cakes Cakes
}

func (reader XMLReader) readDB(file *os.File) (Cakes, error) {
	fmt.Println("bruh")
	return reader.cakes, nil
}
