package main

import "fmt"

const (
	ConferenceName         = "Go Conference"
	ConferenceTotalTickets = 50
)

func ticketInfo(totalTickets int, remainingTickets uint) {
	fmt.Printf("We have total of %v tickets and %v tickets still available\n", totalTickets, remainingTickets)
}

func welcomeMsg(title string, remainingTickets uint) {
	fmt.Println()
	fmt.Printf("----------- %v ----------\n", title)
	fmt.Println()
	fmt.Printf("Welcome to %v booking application\n", title)
	ticketInfo(ConferenceTotalTickets, remainingTickets)
	fmt.Println("Kindly, Get your tickets here to attend.")
	fmt.Println()
}

func main() {
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	welcomeMsg(ConferenceName, remainingTickets)

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, ConferenceName)
}
