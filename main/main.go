package main

import (
	"fmt"

	"rpg-game/game"
	"rpg-game/setup"
)

func main() {
	players := setup.SetupGame()

	// TODO: toss coin

	// Game loop
	game.RunGame(&players, 0)

	fmt.Println(players)
}
