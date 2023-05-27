package main

import "fmt"

func main() {
	var favColor string = "Dodger Blue"
	
	//Not my real age
	 birthYear, age := 1999, 20

	var (
		firstInitial  string = "O"
		secondInitial string = "I"
		job string = "Software Engineer"
	)

	var ageInDays int
	ageInDays = 365 * age
	
	fmt.Println("My favorite color is", favColor)
	fmt.Println("I was born in", birthYear, "and I am", age, "years old")
	fmt.Println("My initials are", firstInitial, secondInitial, "and I am a", job)
	fmt.Println("I am", ageInDays, "days old")
}
