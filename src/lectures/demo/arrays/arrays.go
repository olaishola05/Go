package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func isRoomCleaned(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	isRoomCleaned(rooms)

	fmt.Println("General cleaning ongoing...")

	rooms[2].cleaned = true
	rooms[0].cleaned = true

	isRoomCleaned(rooms)

}
