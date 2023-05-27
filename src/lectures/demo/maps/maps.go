package main

import "fmt"

func main() {
	groceryLists := make(map[string]int)
	groceryLists["eggs"] = 11
	groceryLists["milk"] = 1
	groceryLists["bread"] += 1

	groceryLists["eggs"] += 1
	fmt.Println(groceryLists)

	delete(groceryLists, "milk")
	fmt.Println("Milk deleted, new list:", groceryLists)

	fmt.Println("need", groceryLists["eggs"], "eggs")

	cereal, found := groceryLists["cereal"]
	fmt.Println("Need cereal?")

	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup", cereal, "boxes")
	}

	totalItems := 0

	for _, amount := range groceryLists {
		totalItems += amount
	}

	fmt.Println("There are items", totalItems, "on the grocery lists")
}
