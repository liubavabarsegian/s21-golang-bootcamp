package dbcomparer

import (
	reader "day01/pkg/db_reader"
	"fmt"
	"slices"
)

func Compare(oldCakes reader.Recipes, newCakes reader.Recipes) {
	oldNames := GetNames(oldCakes)

	CompareCakeNames(oldCakes, newCakes)
	// details changes in not new and not removed cakes
	for _, newCake := range newCakes.Cakes {
		if slices.Contains(oldNames, newCake.Name) {
			oldCake := GetGake(newCake.Name, oldCakes)
			CompareTimeChanges(newCake, oldCake)
			CompareIngredientsChanges(newCake, oldCake)
		}
	}
}

func GetNames(recipes reader.Recipes) (names []string) {
	for _, cake := range recipes.Cakes {
		names = append(names, cake.Name)
	}
	return
}

func CompareCakeNames(oldCakes reader.Recipes, newCakes reader.Recipes) {
	oldNames := GetNames(oldCakes)
	newNames := GetNames(newCakes)
	// check for new cakes
	for _, cake := range newCakes.Cakes {
		if !slices.Contains(oldNames, cake.Name) {
			fmt.Printf("ADDED cake \"%s\"\n", cake.Name)
		}
	}

	// check for removed cakes
	for _, cake := range oldCakes.Cakes {
		if !slices.Contains(newNames, cake.Name) {
			fmt.Printf("REMOVED cake \"%s\"\n", cake.Name)
		}
	}
}

func CompareTimeChanges(newCake reader.Cake, oldCake reader.Cake) {
	if newCake.Time != oldCake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", newCake.Name, newCake.Time, oldCake.Time)
	}
}

func GetGake(name string, recipes reader.Recipes) (cake reader.Cake) {
	for _, recipeCake := range recipes.Cakes {
		if recipeCake.Name == name {
			cake = recipeCake
		}
	}
	return
}

func CompareIngredientsChanges(newCake reader.Cake, oldCake reader.Cake) {
	newIngredients, newIngredientsNames := GetIngredients(newCake)
	oldIngredients, oldIngredientsNames := GetIngredients(oldCake)
	// check for new ingredients
	for _, ingredient := range newIngredients {
		if !slices.Contains(oldIngredientsNames, ingredient.Name) {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingredient.Name, newCake.Name)
		}
	}
	// check for removed cakes
	for _, ingredient := range oldIngredients {
		if !slices.Contains(newIngredientsNames, ingredient.Name) {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ingredient.Name, newCake.Name)
		}
	}
	// details changes in not new and not removed imgredients
	for _, newIngredient := range newIngredients {
		if slices.Contains(oldIngredientsNames, newIngredient.Name) {
			oldIngredient := GetIngredient(newIngredient.Name, oldIngredients)
			CompareIngredientUnitsChanges(newIngredient, oldIngredient, newCake.Name)
			CompareIngredientCountsChanges(newIngredient, oldIngredient, newCake.Name)
		}
	}
}

func GetIngredients(cake reader.Cake) (ingredients []reader.Ingredient, names []string) {
	for _, ingredient := range cake.Ingredients {
		names = append(names, ingredient.Name)
		ingredients = append(ingredients, ingredient)
	}
	return
}

func GetIngredient(name string, ingredients []reader.Ingredient) (ingredient reader.Ingredient) {
	for _, value := range ingredients {
		if value.Name == name {
			ingredient = value
			return
		}
	}
	return
}

func CompareIngredientCountsChanges(newIngredient reader.Ingredient, oldIngredient reader.Ingredient, newCakeName string) {
	switch {
	case newIngredient.Count != "" && oldIngredient.Count == "":
		fmt.Printf("ADDED unit count for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCakeName)
	case newIngredient.Count == "" && oldIngredient.Count != "":
		fmt.Printf("REMOVED unit count for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCakeName)
	case newIngredient.Count != oldIngredient.Count:
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
			newIngredient.Name, newCakeName, newIngredient.Count, oldIngredient.Count)
	}
}

func CompareIngredientUnitsChanges(newIngredient reader.Ingredient, oldIngredient reader.Ingredient, newCakeName string) {
	switch {
	case newIngredient.Unit != "" && oldIngredient.Unit == "":
		fmt.Printf("ADDED unit for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCakeName)
	case newIngredient.Unit == "" && oldIngredient.Unit != "":
		fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCakeName)
	case newIngredient.Unit != oldIngredient.Unit:
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
			newIngredient.Name, newCakeName, newIngredient.Unit, oldIngredient.Unit)
	}
}
