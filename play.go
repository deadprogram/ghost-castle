package main

import (
	"strconv"

	"github.com/deadprogram/tinyrogue"
)

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

	game.SetMap(tinyrogue.NewSingleLevelGameMap())

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
	img := game.Images["player"]
	player = NewAdventurer("Sir Scaredy", &img, 5)
	player.ViewRadius = 3
	game.SetPlayer(player)
}

func createGhost() *Ghost {
	img := game.Images["ghost"]
	ghost := NewGhost("Ghost-"+strconv.Itoa(totalGhosts), &img, 60)
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
	level := game.CurrentLevel()
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

func updatePlay() {
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
}
