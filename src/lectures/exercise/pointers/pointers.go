package main

import "fmt"

type SecurityTag bool

type Product struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkOut(slice []Product) {
	for i, _ := range slice {
		deactivate(&slice[i].tag)
	}
}

const (
	Active   = true
	Inactive = false
)

func main() {
	productSlice := []Product{
		{"bags", Active},
		{"shoe", Active},
		{"shoe", Active},
		{"shoe", Active},
	}

	fmt.Println("Initial Items", productSlice)

	deactivate(&productSlice[0].tag)
	fmt.Println("First item deactivated", productSlice)

	checkOut(productSlice)
	fmt.Println("All items deactivated", productSlice)

	activate(&productSlice[3].tag)
	fmt.Println("Item 3 activated", productSlice)
}
