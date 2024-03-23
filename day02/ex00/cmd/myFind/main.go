package main

import (
	"fmt"
	"myFind/config"
	finder "myFind/pkg"
)

func main() {

	arguments, flags, err := config.GetArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	finder.IterateOverEntities(arguments, flags)
}
