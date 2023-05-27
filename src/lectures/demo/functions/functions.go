package main

import "fmt"

func double(number int) int {
	return number * 2
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello, greetings from the greet function")
}

func main() {

	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A bakers dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Another bakers dozen is", anotherBakersDozen)

}
