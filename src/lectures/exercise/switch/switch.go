package main

import "fmt"

func isNewborn(age int) bool {
	return age == 0
}

func isToddler(age int) bool {
	return age >= 1 && age <= 3
}

func isChild(age int) bool {
	return age >= 4 && age <= 12
}

func isTeenager(age int) bool {
	return age >= 13 && age <= 17
}

func main() {

	for i := 0; i <= 100; i++ {
		switch age := i; {
		case isNewborn(age):
			fmt.Println("newborn", age)
		case isToddler(age):
			fmt.Println("toddler", age)
		case isChild(age):
			fmt.Println("child", age)
		case isTeenager(age):
			fmt.Println("teenager", age)
		default:
			fmt.Println("adult", age)
		}
	}
}
