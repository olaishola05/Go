package main

import "fmt"

type Name string

type Player struct {
	name              Name
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "Add", amount, "health ->", player.health)
}

func (player *Player) modifyPlayerHealth(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "damage", amount, "->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount

	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println(player.name, "Add", amount, "Energy ->", player.energy)

}

func (player *Player) modifyPlayerEnergy(amount uint) {
	if player.energy-amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Consume", amount, "->", player.energy)
}

func main() {
	player := Player{
		name:      "kix",
		health:    500,
		maxHealth: 500,
		energy:    2000,
		maxEnergy: 2000,
	}

	player.modifyPlayerHealth(20)
	fmt.Println(player)

	player.addHealth(5)
	fmt.Println(player)

	player.addEnergy(500)
	fmt.Println(player)

	player.modifyPlayerEnergy(1000)
	fmt.Println(player)

	player.modifyPlayerEnergy(1999)
	fmt.Println(player)
}
