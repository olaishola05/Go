package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz1 greater than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 greater than quiz1")
	} else {
		fmt.Println("quiz2 and quiz1 are same scores")
	}

	if avgScore := average(quiz1, quiz2, quiz3); avgScore > 7 {
		fmt.Println("Acceptable grades")
		avgScore += 5
		fmt.Println(avgScore)
	} else {
		fmt.Println("God help the students")
	}
}
