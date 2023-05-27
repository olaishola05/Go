package main

import "fmt"

func main() {
	slice := []string{"hello", "world", "!"}

	for i, element := range slice {
		println(i, element, ":")

		for _, ch := range element {
			fmt.Printf("   %q\n", ch) //%q allow us to print the glyph representation instated bytes
		}
	}
}
