package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printShoppingStats(sum, length int, item Product) {
	fmt.Println("Last item on the List:", item)
	fmt.Println("The sum of shopping list items:", sum)
	fmt.Println("Total items on the shopping list:", length)
}

func shoppingList(lists [4]Product) {
	sum, totalItems := 0, 0

	for i := 0; i < len(lists); i++ {
		item := lists[i]
		sum += item.price

		if item.name != "" {
			totalItems += 1
		}

	}
	printShoppingStats(sum, totalItems, lists[totalItems-1])
}

func main() {
	lists := [4]Product{
		{name: "Orange", price: 5},
		{name: "Bacon", price: 20},
		{name: "Chicken", price: 30},
	}

	shoppingList(lists)

	lists[3] = Product{name: "Turkey", price: 30}

	fmt.Println("Updated List")
	shoppingList(lists)
}
