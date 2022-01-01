package actions

import (
	"math/rand"
	"time"
)

var randNumberSource = rand.NewSource(time.Now().UnixNano())
var randNumberGenerator = rand.New(randNumberSource)
var monsterHealth = MONSTER_HEALTH
var playerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minAttack := PLAYER_ATTACK_MIN_DMG
	maxAttack := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttack = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttack = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}

	damageAmount := getRandomNumber(minAttack, maxAttack)
	monsterHealth -= damageAmount

	return damageAmount
}

func HealPlayer() int {
	minHealValue := PLAYER_HEAL_MIN_VALUE
	maxHealValue := PLAYER_HEAL_MAX_VALUE

	healValue := getRandomNumber(minHealValue, maxHealValue)

	healthDiff := 100 - playerHealth

	if healthDiff >= healValue {
		playerHealth += healValue
		return healValue
	} else {
		playerHealth = 100
		return healthDiff
	}
}

func AttackPlayer() int {
	minAttack := MONSTER_ATTACK_MIN_DMG
	maxAttack := MONSTER_ATTACK_MAX_DMG

	damageAmount := getRandomNumber(minAttack, maxAttack)
	playerHealth -= damageAmount

	return damageAmount
}

func GetHealthAmount() (int, int) {
	return playerHealth, monsterHealth
}

func getRandomNumber(min int, max int) int {
	return randNumberGenerator.Intn(max-min) + min
}
