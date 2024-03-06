package main

import (
	"day01/dbCompare/config"
	dbcomparer "day01/pkg/db_comparer"
	dbreader "day01/pkg/db_reader"
	"fmt"
)

func main() {
	oldDB, newDB, err := config.GetFiles()
	if err != nil {
		fmt.Println(err)
		return
	}

	old_cakes, _ := dbreader.GetCakes(oldDB)
	new_cakes, _ := dbreader.GetCakes(newDB)

	dbcomparer.Compare(old_cakes, new_cakes)
}
