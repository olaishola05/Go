//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func diceRoll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dice, sides := 4, 6
	rolls := 1

	for roll := 1; roll <= rolls; roll++ {
		sum := 0

		for dc := 1; dc <= dice; dc++ {
			rolled := diceRoll(sides)
			sum += rolled
			fmt.Println("Roll #", roll, ", Die #", dc, ":", rolled)
		}

		fmt.Println("Total rolled:", sum)

		switch total := sum; {
		case total == 2 && dice == 2:
			fmt.Println("Snake Eyes!")
		case total == 7:
			fmt.Println("Lucky 7")
		case total%2 == 0:
			fmt.Println("Even")
		case total%2 != 0:
			fmt.Println("Odd")
		}
	}
}
