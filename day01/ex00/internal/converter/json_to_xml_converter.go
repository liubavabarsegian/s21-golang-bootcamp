package converter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Cakes struct {
	Cakes []Cake `json:"cake"`
}

type Cake struct {
	Name        string       `json:"name"`
	Time        string       `json:"time"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name"`
	Count int    `json:"ingredient_count"`
	Unit  string `json:"ingredient_unit"`
}

func JSONtoXML(file *os.File) {
	fmt.Println("JSON TO XML")

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	// we initialize our Users array
	var cakes Cakes

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &cakes)

	fmt.Println(cakes)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(cakes.Cakes); i++ {
		fmt.Println("User Type: " + cakes.Cakes[i].Name)
	}
}
