package main

import "fmt"

func greetings(name string) string {
	fmt.Println("Hello", name, "and welcome to Go land")
	return name
}

func message() string {
	return "Take care of Go Land"
}

func addThreeNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func anyNumber(num int) int {
	return num
}

func twoNumbers() (int, int) {
	return 1, 2
}



func main() {
	greetings("Segun")

	fmt.Println(message())

	a, b := twoNumbers()

	answers := addThreeNumbers(a, b, anyNumber(40))
	fmt.Println("The sum of adding three numbers:", answers)

}
