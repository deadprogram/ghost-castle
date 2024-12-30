package main

import (
	"strconv"

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
		game.Update()

		switch {
		case game.DialogShowing:
			// do nothing, since we might be showing "game over" dialog
			return
		case respawnGhost:
			// all ghosts defeated, respawn them
			respawnGhosts()
		case game.Turn == tinyrogue.GameOver:
			scene = gameOver
			pause = 0
		}
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

func setupGame() {
	game = tinyrogue.NewGame()
	game.UseFOV = true
	game.SetActionSystem(&CombatSystem{})

	loadGameImages()

	gd := tinyrogue.NewGameData(16, 10, 16, 16)
	gd.MinSize = 3
	gd.MaxSize = 6
	gd.MaxRooms = 32
	game.SetData(gd)
}

func startGame() {
	score = 0
	numberGhosts = 1
	totalGhosts = 1

	game.SetMap(tinyrogue.NewGameMap())

	createPlayer()
	ghost := createGhost()

	// set player initial position
	player.MoveTo(game.CurrentLevel().OpenLocation())

	// set monster initial position
	ghost.MoveTo(game.CurrentLevel().OpenLocation())
}

func loadGameImages() {
	game.LoadImage("floor")
	game.LoadImage("wall")
	game.LoadImage("player")
	game.LoadImage("ghost")
}

func createPlayer() {
	player = NewAdventurer("Sir Scaredy", game.Images["player"], 5)
	player.ViewRadius = 2
	game.SetPlayer(player)
}

func createGhost() *Ghost {
	ghost := NewGhost("Ghost-"+strconv.Itoa(totalGhosts), game.Images["ghost"], 60)
	ghost.SetBehavior(tinyrogue.CreatureApproach)
	game.AddCreature(ghost)

	totalGhosts++
	return ghost
}

func removeAllGhosts() {
	for _, c := range game.Creatures {
		if gh, ok := c.(*Ghost); ok {
			removeGhost(gh)
		}
	}
}

func removeGhost(ghost *Ghost) {
	game.RemoveCreature(ghost)
	level := game.Map.CurrentLevel
	level.Block(ghost.GetPosition(), false)
}

func respawnGhosts() {
	respawnDelay++
	if respawnDelay > 120 {
		for i := 0; i < numberGhosts; i++ {
			ghost := createGhost()
			pos := game.CurrentLevel().OpenLocation()
			ghost.MoveTo(pos)
			game.CurrentLevel().Block(pos, true)
		}

		// next level, increase number of ghosts to respawn
		numberGhosts++
		respawnGhost = false
		respawnDelay = 0
	}
}
