package main

import (
	"github.com/deadprogram/tinyrogue"
	"github.com/firefly-zero/firefly-go/firefly"
)

const (
	gameStart = "start"
	gamePlay  = "game"
	gameOver  = "gameover"
)

var (
	scene = gameStart
	pause = 0

	titleFont firefly.Font
	msgFont   firefly.Font

	game *tinyrogue.Game

	player *Adventurer

	// game score, how many ghosts defeated
	score int

	// flag to respawn ghosts
	respawnGhost bool

	// number of ghosts to respawn, increases with each level
	numberGhosts int

	// total number of ghosts in the game, used for naming
	totalGhosts int

	// delay before respawning ghosts
	respawnDelay int
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

func boot() {
	titleFont = firefly.LoadFile("titlefont", nil).Font()
	msgFont = firefly.LoadFile("msgfont", nil).Font()

	setupGame()
}

func update() {
	switch scene {
	case gameStart:
		updateStart()
	case gamePlay:
		updatePlay()
	case gameOver:
		updateGameover()
	}
}

func render() {
	switch scene {
	case gameStart:
		renderStart()
	case gamePlay:
		game.Render()
	case gameOver:
		renderGameover()
	}
}
