package main

import (
	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

// Adventurer represents the player character with combat capabilities
type Adventurer struct {
	*tinyrogue.Player
	*combatant
}

// NewAdventurer creates a new adventurer with specified attributes
func NewAdventurer(name string, image *firefly.Image, speed int) *Adventurer {
	return &Adventurer{
		Player:    tinyrogue.NewPlayer(name, "adventurer", image, speed),
		combatant: NewCombatant(8, 10, 8, "sword"),
	}
}

// Ghost represents an enemy character with combat capabilities
type Ghost struct {
	*tinyrogue.Creature
	*combatant
}

// NewGhost creates a new ghost with specified attributes
func NewGhost(name string, image *firefly.Image, speed int) *Ghost {
	return &Ghost{
		Creature:  tinyrogue.NewCreature(name, "ghost", image, speed),
		combatant: NewCombatant(2, 5, 4, "shriek"),
	}
}
