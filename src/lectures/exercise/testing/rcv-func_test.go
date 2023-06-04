package main

import "testing"

func createPlayer() Player {
	return Player{
		name:      "kix",
		health:    500,
		maxHealth: 500,
		energy:    2000,
		maxEnergy: 2000,
	}
}
func TestHealth(t *testing.T) {
	player := createPlayer()
	player.addHealth(30)

	if player.health > player.maxHealth {
		t.Fatalf("player Health should not be more than: %v", player.maxHealth)
	}

	player.modifyPlayerHealth(player.maxHealth + 1)

	if player.health < 0 {
		t.Fatalf("player health cannot go below %v", 0)
	}

	if player.health > player.maxHealth {
		t.Fatalf("player health should not be greater than player maxHealth %v", player.maxHealth)
	}
}

func TestEnergy(t *testing.T) {
	player := createPlayer()
	player.addEnergy(30)

	if player.energy > player.maxEnergy {
		t.Fatalf("player Energy should not be more than: %v", player.maxEnergy)
	}

	player.modifyPlayerEnergy(player.maxEnergy + 1)

	if player.energy < 0 {
		t.Fatalf("player energy cannot go below %v", 0)
	}

	if player.energy > player.maxEnergy {
		t.Fatalf("player energy should not be greater than player maxEnergy %v", 0)
	}
}
