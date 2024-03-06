package dbcomparer

import (
	reader "day01/pkg/db_reader"
	"fmt"
	"slices"
)

func Compare(old_cakes reader.Cakes, new_cakes reader.Cakes) {
	fmt.Println("i am comparing")

	fmt.Println(old_cakes.Cakes[0].Name, new_cakes.Cakes[0].Name)
	CompareCakeNames(old_cakes, new_cakes)

}

func CompareCakeNames(old_cakes reader.Cakes, new_cakes reader.Cakes) {
	old_names := make([]string, 0)
	for _, value := range old_cakes.Cakes {
		old_names = append(old_names, value.Name)
	}

	new_names := make([]string, 0)
	for _, value := range new_cakes.Cakes {
		new_names = append(new_names, value.Name)
	}

	for _, value := range new_names {
		if !slices.Contains(old_names, value) {
			fmt.Printf("ADDED cake \"%s\"\n", value)
		}
	}

	for _, value := range old_names {
		if !slices.Contains(new_names, value) {
			fmt.Printf("REMOVED cake \"%s\"\n", value)
		}
	}
}
