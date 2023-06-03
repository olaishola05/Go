//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestHealth(t *testing.T) {
	player := Player{
		name:      "kix",
		health:    500,
		maxHealth: 500,
		energy:    2000,
		maxEnergy: 2000,
	}

	player.addHealth(30)

	if player.health > player.maxHealth {
		t.Errorf("player maxHealth should not be more than: %v", player.maxHealth)
	}

	player.addEnergy(500)

	if player.energy > player.maxEnergy {
		t.Errorf("player energy can not be greater than it's maxEnergy: %v", player.maxEnergy)
	}
}

func TestModifyPlayer(t *testing.T) {
	player := Player{
		name:      "kix",
		health:    500,
		maxHealth: 500,
		energy:    2000,
		maxEnergy: 2000,
	}

	player.modifyPlayerHealth(551)

	if player.health < 0 {
		t.Errorf("player health cannot go below %v", 0)
	}
}
