package types

import (
	"fmt"
	"math/rand"
	"rpg-game/utils"
	"time"

	"github.com/manifoldco/promptui"
)

type Health int
type ActionPoints int
type ClassType int

type Attack struct {
	Title           string
	Damage          int
	ActionPointCost ActionPoints
}

type Class struct {
	Title     string
	MaxHealth Health
	Attacks   [4]Attack
}

type Player struct {
	Name         string
	Health       Health
	ActionPoints ActionPoints
	Class        ClassType
	IsDefending  bool
}

type TurnMove int

const (
	MaxActionPoints = 6
	MinPrayHealth   = 5
	MaxPrayHealth   = 15
)

var BasicAttack = Attack{
	Title:           "Basic attack",
	Damage:          1,
	ActionPointCost: 0,
}

const (
	BasicAttackMove TurnMove = iota
	SpecialAttackMove
	Pray
	Defend
)

var Classes = map[ClassType]Class{
	Paladin: {
		"Paladin",
		100,
		[4]Attack{{"Heaven's Judgment", 1, 1}, {"Holy Avenger Slash", 2, 2}, {"Radiant Retribution", 3, 3}, {"Divine Hammerstrike", 4, 4}},
	},
	Wizard: {
		"Wizard",
		80,
		[4]Attack{{"Arcane Nova", 1, 1}, {"Frostbite Lance", 2, 2}, {"Inferno Barrage", 3, 3}, {"Temporal Annihilation", 4, 4}},
	},
	Rogue: {
		"Rogue",
		90,
		[4]Attack{{"Sudden Shadowstrike", 1, 1}, {"Venomous Ambush", 2, 2}, {"Daggerstorm Dance", 3, 3}, {"Evasion Gambit", 4, 4}},
	},
}

func (player *Player) ShowStats() {
	fmt.Printf("Health: %v/%v. Action points: %v/%v\n", player.Health, Classes[player.Class].MaxHealth, player.ActionPoints, MaxActionPoints)
	fmt.Println("---")
}

func (player *Player) Pray() {
	rand.Seed(time.Now().UnixNano())
	restoredHealth := Health(rand.Intn(MaxPrayHealth-MinPrayHealth) + MinPrayHealth)
	player.Health += restoredHealth

	fmt.Printf("%v restored %v health\n", player.Name, restoredHealth)
	// TODO: don't let health be higher than max health
	player.ShowStats()
}

func (player *Player) Defend() {
	player.IsDefending = true
	fmt.Printf("%v is defending", player.Name)
}

// Return true if player won, false if game continues
func (player *Player) UseTurn(opponent *Player) bool {
	utils.ClearScreen()
	fmt.Printf("%v's turn, received 1 Action Point.\n", player.Name)
	player.ShowStats()
	turnMovePrompt := promptui.Select{
		Label: "What is your next move",
		Items: []string{BasicAttackMove.ToString(), SpecialAttackMove.ToString(), Pray.ToString(), Defend.ToString()},
	}

	choice, _, err := turnMovePrompt.Run()

	if err != nil {
		panic("No choice was selected?!")
	}

	switch choice {
	case int(BasicAttackMove):
		player.Attack(opponent, &BasicAttack)
	case int(SpecialAttackMove):
		attacksList := Classes[player.Class].Attacks
		specialAttackMovePrompt := promptui.Select{
			Label: "Select your attack",
			Items: []string{attacksList[0].Title, attacksList[1].Title, attacksList[2].Title, attacksList[3].Title}, // TODO: add back button
			// TODO: add validation for action points
		}

		choice, _, err := specialAttackMovePrompt.Run()

		if err != nil {
			panic("No attack choice was selected?!")
		}

		return player.Attack(opponent, &attacksList[choice])
	case int(Pray):
		player.Pray()
	case int(Defend):
		player.Defend()
	}

	return false
}

// Returns true if attacker won, false if game continues
func (attacker *Player) Attack(attackedPlayer *Player, attackUsed *Attack) bool {
	// Remove action points from attacker
	attacker.ActionPoints = attacker.ActionPoints - attackUsed.ActionPointCost

	// Remove health from attacked player
	attackedPlayer.Health = attackedPlayer.Health - Health(attackUsed.Damage)

	if attackedPlayer.IsDefending {
		rand.Seed(time.Now().UnixNano())
		attackedPlayer.Health += Health(rand.Intn(6-1) + 1)
		attackedPlayer.IsDefending = false

	}

	// Check if attacked player is dead
	return attackedPlayer.Health <= 0
}

const (
	Paladin ClassType = iota
	Wizard
	Rogue
)

func (c ClassType) ToString() string {
	switch c {
	case Paladin:
		return "Paladin"
	case Wizard:
		return "Wizard"
	case Rogue:
		return "Rogue"
	default:
		panic("Class does not exist")
	}
}

func (m TurnMove) ToString() string {
	switch m {
	case BasicAttackMove:
		return "Basic attack"
	case SpecialAttackMove:
		return "Special attack"
	case Pray:
		return "Pray"
	case Defend:
		return "Defend"
	default:
		panic("Move does not exist")
	}
}
