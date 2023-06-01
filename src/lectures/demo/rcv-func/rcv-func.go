package main

import "fmt"

type Space struct {
	ocupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupyAspace(parkingLot *ParkingLot, spaceNumber int) {
	parkingLot.spaces[spaceNumber-1].ocupied = true
}

func (parkingLot *ParkingLot) occupyAspace(spaceNumber int) {
	parkingLot.spaces[spaceNumber-1].ocupied = true
}

func (parkingLot *ParkingLot) vacateSpace(spaceNumber int) {
	parkingLot.spaces[spaceNumber-1].ocupied = false
}

func main() {
	parkingLot := ParkingLot{spaces: make([]Space, 6)}
	fmt.Println("Intial parking space", parkingLot.spaces)
	parkingLot.occupyAspace(1)
	parkingLot.occupyAspace(5)
	fmt.Println("After occupied", parkingLot.spaces)

	parkingLot.vacateSpace(1)
	fmt.Println("Vacating a space", parkingLot.spaces)

	occupyAspace(&parkingLot, 3)
	fmt.Println("Occupied using pointer func", parkingLot.spaces)

}
