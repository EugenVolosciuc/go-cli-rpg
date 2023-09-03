package setup

import (
	"fmt"
	"rpg-game/types"

	"github.com/manifoldco/promptui"
)

const defaultActionPoints = 2

func setupPlayer(playerNumber int) types.Player {
	// Setup name
	fmt.Printf("Player %v, please enter your name:\n", playerNumber+1)

	namePrompt := promptui.Prompt{
		Label: "Name: ",
	}

	name, err := namePrompt.Run()

	for err != nil {
		fmt.Println("Please write your name")
		name, _ = namePrompt.Run()
	}

	// Setup class
	fmt.Printf("%v, please choose a class\n", name)

	classPrompt := promptui.Select{
		Label: "Class: ",
		Items: []string{types.Paladin.ToString(), types.Wizard.ToString(), types.Rogue.ToString()},
	}

	class, _, err := classPrompt.Run()

	for err != nil {
		fmt.Println("Please select your class")
		class, _, _ = classPrompt.Run()
	}

	return types.Player{
		Name:         name,
		Health:       types.Classes[types.ClassType(class)].MaxHealth,
		ActionPoints: defaultActionPoints,
		Class:        types.ClassType(class),
	}
}

func SetupGame() []types.Player {
	const numberOfPlayers = 2
	fmt.Println("Welcome to Go RPG!")
	fmt.Println("------------------")

	var players []types.Player = make([]types.Player, numberOfPlayers)

	for i := 0; i < numberOfPlayers; i++ {
		player := setupPlayer(i)

		players[i] = player
	}

	return players
}
