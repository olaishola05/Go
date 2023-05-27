package main

import "fmt"

func main() {
	var myName string = "Oladipupo Ishola"

	var (
		myAge    int    = 20
		myHeight float32 = 1.7
	)

	var part1, part2, _ = "Hello", "World", "!"

	data, err := "some random data", "some random error"

	newData, err := "some new data", "some new error"

	fmt.Println(newData, err)

	fmt.Println(data, err)

	fmt.Println(part1, part2)

	fmt.Println("Hello,", myName)

	fmt.Println("I am", myAge, "years old", "and", myHeight, "meters tall")
}
