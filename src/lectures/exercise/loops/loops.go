package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		switch num := i; {
		case num % 3 == 0 && num % 5 == 0:
			fmt.Println("FizzBuzz")
		case num % 3 == 0:
			fmt.Println("Fizz")
		case num % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
