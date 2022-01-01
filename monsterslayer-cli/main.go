package main

import (
	"github/mikemando/monsterslayer/actions"
	"github/mikemando/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	// PLAYER || MONSTER || "" (neither) wins game round
	winner := ""

	for winner == "" {
		winner = executeGame()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeGame() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvaliableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerAttackDamage int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDamage = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDamage = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmount()

	roundData := interaction.RoundData{
		Action:              userChoice,
		PlayerHealth:        playerHealth,
		MonsterHealth:       monsterHealth,
		PlayerAttackingDmg:  playerAttackDamage,
		PlayerHealValue:     playerHealValue,
		MonsterAttackingDmg: monsterAttackDmg,
	}

	interaction.RoundInfo(&roundData)

	gameRounds = append(gameRounds, roundData)

	if playerHealth <= 0 {
		return "MONSTER"
	} else if monsterHealth <= 0 {
		return "PLAYER"
	}
	return ""
}

func endGame(winner string) {
	interaction.GameWinner(winner)
	interaction.LogFile(&gameRounds)
}
