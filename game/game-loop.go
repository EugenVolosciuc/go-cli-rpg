package game

import (
	"fmt"
	"rpg-game/types"
)

func RunGame(players *[]types.Player, firstPlayerToMove int) {
	var playerToPlay *types.Player = &(*players)[firstPlayerToMove]
	var playerToPlayIndex = firstPlayerToMove
	var opponent *types.Player = &(*players)[firstPlayerToMove%2]

	var winner *types.Player

	fmt.Println("-------------------")
	fmt.Println("Let the game begin!")
	fmt.Println("-------------------")

	for winner == nil {
		// TODO: fix game end
		var playerWon = playerToPlay.UseTurn(opponent)

		if playerWon {
			winner = playerToPlay
			fmt.Printf("%v won!", winner.Name)
		} else {
			opponent = playerToPlay
			playerToPlayIndex = (playerToPlayIndex + 1) % 2
			playerToPlay = &(*players)[playerToPlayIndex]
		}
	}
}
