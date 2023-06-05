package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("Cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")

	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}

	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Masala"), Salad("Cream Salad")}

	PrepareDishes(dishes)
}
