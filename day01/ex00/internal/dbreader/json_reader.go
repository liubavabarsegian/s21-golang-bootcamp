package dbreader

import (
	"fmt"
	"os"
)

type JSONReader struct {
	cakes Cakes
}

func (reader JSONReader) readDB(file *os.File) (Cakes, error) {
	fmt.Println("bruh")
	return reader.cakes, nil
}
