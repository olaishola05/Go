package main

import "fmt"

type ModelName string

type Lift int

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type LiftPicker interface {
	PickVehicleLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motorcycle) PickVehicleLift() Lift {
	return SmallLift
}

func (c Car) PickVehicleLift() Lift {
	return StandardLift
}

func (t Truck) PickVehicleLift() Lift {
	return LargeLift
}

func directToLift(p LiftPicker) {
	switch p.PickVehicleLift() {
	case SmallLift:
		fmt.Printf("Send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("Send %v to standard lift\n", p)
	case LargeLift:
		fmt.Printf("Send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Lambogini")
	truck := Truck("Man Diesel")
	motorcycle := Motorcycle("Yamaha")

	directToLift(car)
	directToLift(truck)
	directToLift(motorcycle)
}
