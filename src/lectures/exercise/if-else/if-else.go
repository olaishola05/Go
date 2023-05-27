package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekdays(day int) bool {
	return day <= 4
}

func main() {
	
	day, role := Saturday, Member

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekdays(day) {
		accessGranted()
	} else if role == Member && weekdays(day) {
		accessGranted()
	} else if role == Guest && (day == Monday || day == Wednesday || day == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}
}
