package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Air = iota
	Ground
)

type BeltSize int
type Shipping int

func (s Shipping) String() string {
	return []string{"Small", "Medium", "Large"}[s]
}

func (b BeltSize) String() string {
	return []string{"Ground", "Air"}[b]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WareHouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WareHouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

func main() {
	mail := SpamMail{400000}
	automate(&mail)
}
